package pool_helper

import "sync"

type Client interface {
	Close() error
}

type NewClient func() (Client, error)

type Pool interface {
	Get() (Client, error)
	Put(Client) error
	Max() int
	Size() int
}

func NewPool(initial, max int, new_client NewClient) (Pool, error) {
	clients := make(chan Client, max)

	for i := 0; i < initial; i++ {
		if cli, err := new_client(); err != nil {
			return nil, err
		} else {
			clients <- cli
		}
	}

	return &pool{
		max:        max,
		clients:    make(chan Client),
		new_client: new_client,
	}, nil
}

type pool struct {
	max        int
	clients    chan Client
	new_client NewClient
	mtx        sync.Mutex
}

func (p *pool) Get() (Client, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	select {
	case cli := <-p.clients:
		return cli, nil
	default:
		return p.new_client()
	}
}

func (p *pool) Put(c Client) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	select {
	case p.clients <- c:
		return nil
	default:
		return c.Close()
	}
}

func (p *pool) size() int {
	return len(p.clients)
}

func (p *pool) Size() int {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	return p.size()
}

func (p *pool) Max() int {
	return p.max
}
