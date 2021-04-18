package ami

import (
	"context"
	"errors"
	"sync"
)

//NewPool return an AMI connection Pool
// connection pools allows to have multple AMI connections with asterisk
// it is useful for async requests to asterisk (i.e. api requests)
func NewPool(ctx context.Context, address string, user, secret, events string) (*Pool, error) {
	return &Pool{
		MinConections: 0,
		MaxConections: 10,

		ctx:      ctx,
		address:  address,
		username: user,
		secret:   secret,
		events:   events,
		mutex:    new(sync.Mutex),
		active:   make(map[*Socket]bool),
		idle:     make([]*Socket, 0),
		closed:   true,
	}, nil
}

//Pool is the struct that mantains all the asterisk connections open
type Pool struct {
	MinConections int //initial ami connections
	MaxConections int //max connections allowed
	address       string
	username      string
	secret        string
	events        string
	ctx           context.Context
	idle          []*Socket        //list of idle asterisk connections
	active        map[*Socket]bool //index of active(in use) connections
	mutex         *sync.Mutex
	closed        bool
}

//Connect connects with every ami socket in the pool
func (p *Pool) Connect() error {
	p.closed = false
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for i := 0; i < p.MinConections; i++ {
		socket, err := p.newConnection()
		if err != nil {
			return err
		}
		p.idle = append(p.idle, socket)
	}
	return nil
}

//newConnection a new socket with asterisk, then login to the ami interface
func (p *Pool) newConnection() (*Socket, error) {

	totalSessions := len(p.active) + len(p.idle)
	if totalSessions > p.MaxConections {
		return nil, errors.New("Max allowed connections reached. Increase the max allowed connections by setting pool.MaxConections to a higher value")
	}

	socket, err := NewSocket(p.ctx, p.address)
	if err != nil {
		return nil, err
	}
	if _, err := Connect(p.ctx, socket); err != nil {
		return nil, err
	}
	if err := Login(p.ctx, socket, p.username, p.secret, p.events, ""); err != nil {
		return nil, err
	}
	return socket, nil
}

//CloseAll closes all AMI connections and shutdown pool
func (p *Pool) CloseAll() error {
	p.closed = true

	for _, s := range p.idle {
		Logoff(p.ctx, s, "")
		s.Close(p.ctx)
	}

	for s := range p.active {
		Logoff(p.ctx, s, "")
		s.Close(p.ctx)
	}

	return nil
}

//GetSocket return a connected AMI session
// if no idle connection is found, this function creates a new socket
func (p *Pool) GetSocket() (*Socket, error) {
	if p.closed {
		return nil, errors.New("Pool closed")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var s *Socket
	var err error
	if len(p.idle) == 0 {
		s, err = p.newConnection()
		if err != nil {
			return nil, err
		}
	} else {
		s = p.idle[0]
		p.idle = p.idle[1:]
	}

	p.active[s] = true

	return s, nil
}

//Close give back a socket to the pool
// if the idle connections are greater than MinConections, then this connection is closed
// the "force" param is set to true, the connection to asterisk will be closed
func (p *Pool) Close(s *Socket, force bool) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.active, s)
	totalSessions := len(p.active) + len(p.idle)
	if totalSessions >= p.MinConections || force {
		s.Close(p.ctx)
		return nil
	}

	p.idle = append(p.idle, s)

	return nil
}
