package ioc_test

import (
	"testing"
	"unsafe"

	"github.com/ronbb/ioc"
)

type S struct {
	A int
}

func TestSingleton(t *testing.T) {
	c1 := ioc.NewContainer()
	s1 := &S{
		A: 1,
	}
	c1.Singleton(&s1)
	var s2, s3 *S
	println(uintptr(unsafe.Pointer(&s2)))
	c1.Make(&s2)
	c1.Make(&s3)

	println(uintptr(unsafe.Pointer(&s2)))

	if uintptr(unsafe.Pointer(s1)) != uintptr(unsafe.Pointer(s2)) {
		t.Error("[Singleton] s1 and s2 are not same address")
	}

	if uintptr(unsafe.Pointer(s2)) != uintptr(unsafe.Pointer(s3)) {
		t.Error("[Singleton] s2 and s3 are not same address")
	}
}

func TestFactory(t *testing.T) {
	c2 := ioc.NewContainer()
	c2.Factory(func() *S {
		s := &S{}
		return s
	})
	var s4, s5 *S
	c2.Make(&s4)
	c2.Make(&s5)

	if uintptr(unsafe.Pointer(s4)) == uintptr(unsafe.Pointer(s5)) {
		t.Errorf("[Factory] same address %p %p", s4, s5)
	}

	var s8, s9 *S
	c2.Make(func(s *S) {
		s8 = s
	})
	c2.Make(func(s *S) {
		s9 = s
	})

	if uintptr(unsafe.Pointer(s8)) == uintptr(unsafe.Pointer(s9)) {
		t.Errorf("[Factory] same address %p %p", s8, s9)
	}
}

func TestLazy(t *testing.T) {
	c3 := ioc.NewContainer()
	c3.Lazy(func() *S {
		s := &S{}
		return s
	})
	var s6, s7 *S
	c3.Make(&s6)
	c3.Make(&s7)

	if uintptr(unsafe.Pointer(s6)) != uintptr(unsafe.Pointer(s7)) {
		t.Errorf("[Lazy] different address %p %p", s6, s7)
	}

	var s8, s9 *S
	c3.Make(func(s *S) {
		s8 = s
	})
	c3.Make(func(s *S) {
		s9 = s
	})

	if uintptr(unsafe.Pointer(s8)) != uintptr(unsafe.Pointer(s9)) {
		t.Errorf("[Lazy] different address %p %p", s8, s9)
	}
}
