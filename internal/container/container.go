package container

import (
	"reflect"
)

// Container ioc container
type Container interface {
	Singleton(inst interface{})
	Lazy(inst interface{})
	Factory(inst interface{})
	Reset()
	Make(interface{})
}

// New create an ioc container
func New() Container {
	return &container{
		bindings: make(map[reflect.Type]binding),
	}
}

type container struct {
	bindings map[reflect.Type]binding
	// TODO
}

func (c *container) arguments(function interface{}) []reflect.Value {
	funcType := reflect.TypeOf(function)
	argsCount := funcType.NumIn()
	args := make([]reflect.Value, argsCount)

	for i := 0; i < argsCount; i++ {
		abstraction := funcType.In(i)

		var instance interface{}

		if concrete, ok := c.bindings[abstraction]; ok {
			instance = concrete.resolve(c)
		} else {
			panic("no concrete found for the abstraction: " + abstraction.String())
		}

		args[i] = reflect.ValueOf(instance)
	}

	return args
}

func (c *container) bind(resolver interface{}, bindingType bindingType) {
	resolverType := reflect.TypeOf(resolver)
	if resolverType == nil {
		panic("resolver should not be nil")
	}

	switch resolverType.Kind() {
	case reflect.Ptr:
		if bindingType != singleton {
			panic("ptr resolver can only bind a singleton")
		}
		c.bindPointer(resolver, resolverType)
		break
	case reflect.Func:
		c.bindFunction(resolver, resolverType, bindingType)
		break
	default:
		panic("unavailable resolver type")
	}
}

func (c *container) bindPointer(resolver interface{}, resolverType reflect.Type) {
	c.bindings[resolverType.Elem()] = &singletonBinding{
		instance: extract(resolver).Interface(),
	}
}

func (c *container) bindFunction(resolver interface{}, resolverType reflect.Type, bindingType bindingType) {
	numOut := resolverType.NumOut()
	if numOut < 1 {
		panic("the number of out of resolver must be at least one")
	}

	switch bindingType {
	case singleton:
		outs := make([]reflect.Value, numOut, numOut)
		for i := 0; i < numOut; i++ {
			c.bindings[resolverType.Out(i)] = &singletonBinding{
				instance: outs[i].Interface(),
			}
		}
		break
	case factory:
		for i := 0; i < numOut; i++ {
			c.bindings[resolverType.Out(i)] = &factoryBinding{
				outIndex: i,
				resolver: resolver,
			}
		}
		break
	case lazy:
		for i := 0; i < numOut; i++ {
			c.bindings[resolverType.Out(i)] = &lazyBinding{
				resolved: false,
				outIndex: i,
				resolver: resolver,
			}
		}
		break
	default:
		panic("unknown binding type")
	}
}

func (c *container) make(receiver interface{}) {
	receiverType := reflect.TypeOf(receiver)
	if receiverType == nil {
		panic("receiver should not be nil")
	}

	switch receiverType.Kind() {
	case reflect.Ptr:
		c.makePointer(receiver, receiverType)
		break
	case reflect.Func:
		c.makeFunction(receiver, receiverType)
		break
	default:
		panic("unavailable receiver type")
	}

}

func (c *container) makePointer(receiver interface{}, receiverType reflect.Type) {
	abstraction := receiverType.Elem()

	if concrete, ok := c.bindings[abstraction]; ok {
		instance := concrete.resolve(c)
		reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(instance))
		return
	}

	panic("no concrete found for the abstraction " + abstraction.String())
}

func (c *container) makeFunction(receiver interface{}, receiverType reflect.Type) {
	arguments := c.arguments(receiver)
	reflect.ValueOf(receiver).Call(arguments)
	return
}
