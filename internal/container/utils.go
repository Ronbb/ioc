package container

import "reflect"

func invoke(c *container, function interface{}) []reflect.Value {
	return reflect.ValueOf(function).Call(c.arguments(function))
}

func extract(ptr interface{}) reflect.Value {
	return reflect.ValueOf(ptr).Elem()
}
