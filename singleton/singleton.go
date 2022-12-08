package singleton

import (
	"math/rand"
	"sync"
)

type singleton struct {
	randInt int
}

type Singleton interface {
	Get() int
}

func (s *singleton) Get() int {
	return s.randInt
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{
			randInt: rand.Int(),
		}
	})
	return instance
}
