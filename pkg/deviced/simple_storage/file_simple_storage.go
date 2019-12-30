package metathings_deviced_simple_storage

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	file_helper "github.com/nayotta/metathings/pkg/common/file"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
)

type FileSimpleStorageOption struct {
	Home string
}

func NewFileSimpleStorageOption() *FileSimpleStorageOption {
	return &FileSimpleStorageOption{}
}

type FileSimpleStorage struct {
	opt    *FileSimpleStorageOption
	logger log.FieldLogger
}

// Copy from ioutil.tempfile
// Random number state.
// We generate random temporary file names so that there's a good
// chance the file doesn't exist yet - keeps the number of tries in
// TempFile to a minimum.
var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextRandom() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func (fss *FileSimpleStorage) join_path(obj *Object) string {
	return path.Join(fss.opt.Home, obj.Device, obj.FullName())
}

func (fss *FileSimpleStorage) join_temp_path(obj *Object) string {
	fn := obj.FullName()
	fn_dir := path.Dir(fn)
	fn_base := path.Base(fn)
	return path.Join(fss.opt.Home, obj.Device, "tmp", fn_dir, fn_base+"."+nextRandom())
}

func (fss *FileSimpleStorage) is_empty(dev *storage.Device, obj *Object) (bool, error) {
	p := fss.join_path(obj)
	f, err := os.Open(path.Dir(p))
	if err != nil {
		return false, err
	}

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, nil
}

func (fss *FileSimpleStorage) etag(obj *Object) string {
	s := fmt.Sprintf("%v#%v#%v", fss.join_path(obj), obj.Length, obj.LastModified.UnixNano())
	h := sha1.Sum([]byte(s))
	return base64.StdEncoding.EncodeToString([]byte(h[:]))
}

func (fss *FileSimpleStorage) new_object(device, prefix, name string, length int64, last_modified time.Time) *Object {
	obj := new_object(device, prefix, name, length, "\"\"", last_modified)
	obj.Etag = fss.etag(obj)
	return obj
}

func (fss *FileSimpleStorage) PutObject(obj *Object, reader io.Reader) error {
	p := fss.join_path(obj)

	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(p, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	slice := make([]byte, 4096)
	for n, err := reader.Read(slice); n > 0 && err == nil; n, err = reader.Read(slice) {
		if n, err = f.Write(slice[:n]); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (fss *FileSimpleStorage) PutObjectAsync(obj *Object, opt *PutObjectAsyncOption) (*file_helper.FileSyncer, error) {
	p := fss.join_path(obj)
	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return nil, err
	}

	cp := fss.join_temp_path(obj)
	err = os.MkdirAll(path.Dir(cp), os.ModePerm)
	if err != nil {
		return nil, err
	}

	fs, err := file_helper.NewFileSyncer(
		file_helper.SetPath(p),
		file_helper.SetSize(obj.Length),
		file_helper.SetSha1Hash(opt.SHA1),
		file_helper.SetChunkSize(opt.ChunkSize),
		file_helper.SetCachePath(cp),
	)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func (fss *FileSimpleStorage) RemoveObject(obj *Object) error {
	p := fss.join_path(obj)
	pre := path.Dir(p)

	err := os.Remove(p)
	if err != nil {
		return err
	}

	dir, err := os.Open(pre)
	if err != nil {
		return err
	}

	_, err = dir.Readdir(1)
	if err == io.EOF {
		err = os.Remove(pre)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fss *FileSimpleStorage) RenameObject(src, dst *Object) error {
	dst.Device = src.Device
	psrc := fss.join_path(src)
	pdst := fss.join_path(dst)
	predst := path.Dir(pdst)

	err := os.MkdirAll(predst, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Rename(psrc, pdst)
	if err != nil {
		return err
	}

	return nil
}

func (fss *FileSimpleStorage) GetObject(x *Object) (y *Object, err error) {
	p := fss.join_path(x)

	fi, err := os.Stat(p)
	if err != nil {
		return
	}

	if fi.IsDir() {
		y = fss.new_object(x.Device, p, "", fi.Size(), fi.ModTime())
	} else {
		y = fss.new_object(x.Device, path.Dir(p), path.Base(p), fi.Size(), fi.ModTime())
	}

	return
}

func (fss *FileSimpleStorage) GetObjectContent(obj *Object) (chan []byte, error) {
	p := fss.join_path(obj)

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	ch := make(chan []byte)

	go func() {
		slice := make([]byte, 512)
		for {
			n, err := f.Read(slice)
			if err != nil || n == 0 {
				break
			}
			ch <- slice[:n]
		}
		defer close(ch)
	}()

	return ch, nil
}

func (fss *FileSimpleStorage) list_objects(obj *Object, recursive bool, depth int) ([]*Object, error) {
	if recursive && depth == 0 {
		return nil, nil
	}

	obj.Name = ""
	p := fss.join_path(obj)

	fs, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	var objs []*Object
	for _, f := range fs {
		var new_obj *Object
		if f.IsDir() {
			new_obj = fss.new_object(obj.Device, path.Join(obj.Prefix, f.Name()), "", f.Size(), f.ModTime())
			objs = append(objs, new_obj)

			if recursive {
				sub_objs, err := fss.list_objects(new_obj, recursive, depth-1)
				if err != nil {
					return nil, err
				}
				objs = append(objs, sub_objs...)
			}
		} else {
			new_obj = fss.new_object(obj.Device, obj.Prefix, f.Name(), f.Size(), f.ModTime())
			objs = append(objs, new_obj)
		}

	}

	return objs, nil
}

func (fss *FileSimpleStorage) ListObjects(obj *Object, opt *ListObjectsOption) ([]*Object, error) {
	return fss.list_objects(obj, opt.Recursive, opt.Depth)
}

func new_file_simple_storage(args ...interface{}) (SimpleStorage, error) {
	var logger log.FieldLogger
	opt := NewFileSimpleStorageOption()

	err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"home":   opt_helper.ToString(&opt.Home),
		"logger": opt_helper.ToLogger(&logger),
	}, opt_helper.SetSkip(true))(args...)
	if err != nil {
		return nil, err
	}

	return &FileSimpleStorage{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_simple_storage_factory("file", new_file_simple_storage)
}
