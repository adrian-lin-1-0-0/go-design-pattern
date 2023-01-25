package objectpool

import (
	"errors"
	"sync"
)

var NoIdle = errors.New("No idle")

type IPool interface {
	Get() (IPoolObject, error)
}

type IPoolObject interface {
	GetID() string
}

type IPoolObjectFactory interface {
	Create() IPoolObject
}

type pool struct {
	idle     []IPoolObject
	active   []IPoolObject
	capacity int
	mux      sync.Mutex
	factory  IPoolObjectFactory
}

func NewPool(capacity int, factory IPoolObjectFactory) IPool {
	return &pool{
		capacity: capacity,
		factory:  factory,
	}
}

func (p *pool) Get() (IPoolObject, error) {
	p.mux.Lock()

	var obj IPoolObject
	var err error

	if len(p.idle) == 0 {
		if len(p.active) < p.capacity {
			obj = p.factory.Create()
			p.active = append(p.active, obj)
			goto RERUEN
		}

		err = NoIdle
		goto RERUEN
	}

	obj, p.idle = p.idle[0], p.idle[1:]
	p.active = append(p.active, obj)
	goto RERUEN

RERUEN:
	p.mux.Unlock()
	return obj, err
}
