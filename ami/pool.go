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
		LowWater:  0,
		HighWater: 10,

		ctx:      ctx,
		address:  address,
		username: user,
		secret:   secret,
		events:   events,
		mutex:    new(sync.Mutex),
		active:   make(map[*Socket]bool),
		idle:     make([]*Socket, 0),
	}, nil
}

//Pool is the struct that mantains all the asterisk connections open
//TODO: implement a keepalive go rutine that checks idle connections status
type Pool struct {
	LowWater  int //initial ami connections
	HighWater int //max connections allowed
	address   string
	username  string
	secret    string
	events    string
	ctx       context.Context
	idle      []*Socket        //list of idle asterisk connections
	active    map[*Socket]bool //index of active(in use) connections
	mutex     *sync.Mutex
}

//Connect connects with every ami socket in the pool
func (p *Pool) Connect() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for i := 0; i < p.LowWater; i++ {
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
	if totalSessions > p.HighWater {
		return nil, errors.New("Max allowed connections reached. Increase the max allowed connections by setting pool.HighWater to a higher value")
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

//Close closes all AMI connections
func (p *Pool) Close() error {
	return nil
}

//GetSocket return a connected AMI session
func (p *Pool) GetSocket() (*Socket, error) {
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
		//TODO: verify that connection is still alive
		s = p.idle[0]
		p.idle = p.idle[1:]
	}

	p.active[s] = true

	return s, nil
}

//FreeSocket give back a socket to the pool
// if the idle connections are greater than LowWater, then this connection is closed
// the "close" param is set to true, the connection to asterisk will be closed
func (p *Pool) FreeSocket(s *Socket, close bool) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.active, s)
	totalSessions := len(p.active) + len(p.idle)
	if totalSessions >= p.LowWater || close {
		Logoff(p.ctx, s, "")
		s.Close(p.ctx)
		return nil
	}

	p.idle = append(p.idle, s)

	return nil
}
