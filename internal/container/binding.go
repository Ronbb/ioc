package container

import (
	"sync"
)

type bindingType int32

const (
	factory bindingType = iota
	singleton
	lazy
)

type binding interface {
	resolve(c *container) interface{}
}

type singletonBinding struct {
	instance interface{}
}

func (b *singletonBinding) resolve(c *container) interface{} {
	return b.instance
}

type lazyBinding struct {
	mutex    sync.Mutex
	resolved bool
	outIndex int
	resolver interface{}
	instance interface{}
}

func (b *lazyBinding) resolve(c *container) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if !b.resolved {
		instance := invoke(c, b.resolver)[b.outIndex].Interface()
		b.instance = instance
		b.resolved = true
		return instance
	}
	return b.instance
}

type factoryBinding struct {
	outIndex int
	resolver interface{}
}

func (b *factoryBinding) resolve(c *container) interface{} {
	return invoke(c, b.resolver)[b.outIndex].Interface()
}
