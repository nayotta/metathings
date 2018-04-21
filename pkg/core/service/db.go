package metathings_core_service

import "errors"

var (
	NotFound = errors.New("not found")
)

type Core struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	ProjectId string `db:"project_id"`
	OwnerId   string `db:"owner_id"`
	State     string `db:"state"`
}

type Storage interface {
	CreateCore(*Core) (*Core, error)
	DeleteCore(*Core) error
	PatchCore(*Core) (*Core, error)
	GetCore(*Core) (*Core, error)
	ListCores(*Core) ([]*Core, error)
}

type memStorage struct {
	mem map[string]*Core
}

func (s *memStorage) CreateCore(core *Core) (*Core, error) {
	s.mem[core.Id] = core
	return core, nil
}

func (s *memStorage) DeleteCore(core *Core) error {
	if _, ok := s.mem[core.Id]; !ok {
		return NotFound
	}
	delete(s.mem, core.Id)
	return nil
}

func (s *memStorage) PatchCore(core *Core) (*Core, error) {
	if c, ok := s.mem[core.Id]; ok {
		if core.Name != "" {
			c.Name = core.Name
		}

		if core.State != "" {
			c.State = core.State
		}

		return c, nil
	}

	return nil, NotFound
}

func (s *memStorage) GetCore(core *Core) (*Core, error) {
	if c, ok := s.mem[core.Id]; ok {
		return c, nil
	}
	return nil, NotFound
}

func (s *memStorage) ListCores(core *Core) ([]*Core, error) {
	cores := []*Core{}
	for _, c := range s.mem {
		cores = append(cores, c)
	}
	return cores, nil
}

func newMemStorage() *memStorage {
	s := &memStorage{
		mem: make(map[string]*Core),
	}
	return s
}

func NewStorage() (Storage, error) {
	return newMemStorage(), nil
}