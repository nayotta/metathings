package file_helper

const (
	DEFAULT_CHUNK_SIZE int64 = 512 * 1024
)

type SetFileSyncerOption func(*FileSyncerOption)

type FileSyncerOption struct {
	path       string
	size       int64
	cache_path string
	sha1_hash  string
	chunk_size int64
}

func SetPath(path string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.path = path
	}
}

func SetSize(size int64) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.size = size
	}
}

func SetCachePath(path string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.cache_path = path
	}
}

func SetSha1Hash(hash string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.sha1_hash = hash
	}
}

func SetChunkSize(size int64) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.chunk_size = size
	}
}

func NewFileSyncerOption() *FileSyncerOption {
	o := &FileSyncerOption{}
	o.chunk_size = DEFAULT_CHUNK_SIZE
	return o
}

type FileSyncer interface {
	Next(count int) (offsets []int64, err error)
	Sync(offset int64, data []byte, size int) (err error)
	Close() error
}
