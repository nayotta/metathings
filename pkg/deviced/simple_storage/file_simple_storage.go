package metathings_deviced_simple_storage

import (
	"io"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
)

type FileSimpleStorageOption struct {
	Home string
}

func NewFileSimpleStorageOption() *FileSimpleStorageOption {
	return &FileSimpleStorageOption{}
}

// TODO(Peer): let Metadata and Etag work.
type FileSimpleStorage struct {
	opt    *FileSimpleStorageOption
	logger log.FieldLogger
}

func (fss *FileSimpleStorage) join_path(dev *storage.Device, obj *Object) string {
	return path.Join(fss.opt.Home, *dev.Id, obj.FullName())
}

func (fss *FileSimpleStorage) is_empty(dev *storage.Device, obj *Object) (bool, error) {
	p := fss.join_path(dev, obj)
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

func (fss *FileSimpleStorage) PutObject(dev *storage.Device, obj *Object, reader io.Reader) error {
	p := fss.join_path(dev, obj)

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

func (fss *FileSimpleStorage) RemoveObject(dev *storage.Device, obj *Object) error {
	p := fss.join_path(dev, obj)
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

func (fss *FileSimpleStorage) RenameObject(dev *storage.Device, src, dst *Object) error {
	psrc := fss.join_path(dev, src)
	pdst := fss.join_path(dev, dst)
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

func (fss *FileSimpleStorage) GetObject(dev *storage.Device, obj *Object) (*Object, error) {
	p := fss.join_path(dev, obj)

	fi, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	new_obj := new_object(*dev.Id, obj.Prefix, obj.Name, fi.Size(), "\"\"", fi.ModTime(), obj.Metadata)

	return new_obj, nil
}

func (fss *FileSimpleStorage) GetObjectContent(dev *storage.Device, obj *Object) (chan []byte, error) {
	p := fss.join_path(dev, obj)

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

func (fss *FileSimpleStorage) ListObjects(dev *storage.Device, obj *Object) ([]*Object, error) {
	obj.Name = ""
	p := fss.join_path(dev, obj)

	fs, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	var objs []*Object
	for _, f := range fs {
		var new_obj *Object
		if f.IsDir() {
			new_obj = new_object(*dev.Id, path.Join(obj.Prefix, f.Name()), "", f.Size(), "\"\"", f.ModTime(), nil)
		} else {
			new_obj = new_object(*dev.Id, obj.Prefix, f.Name(), f.Size(), "\"\"", f.ModTime(), nil)
		}
		objs = append(objs, new_obj)
	}

	return objs, nil
}

func new_file_simple_storage(args ...interface{}) (SimpleStorage, error) {
	var ok bool
	var logger log.FieldLogger
	opt := &FileSimpleStorageOption{}

	err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"home": func(key string, val interface{}) error {
			opt.Home, ok = val.(string)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
		"logger": func(key string, val interface{}) error {
			logger, ok = val.(log.FieldLogger)
			if !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args...)
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
