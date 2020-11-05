package ioc

import "github.com/ronbb/ioc/internal/container"

var (
	instance Container = nil
)

func init() {
	instance = NewContainer()
}

// NewContainer create a new ioc container
func NewContainer() Container {
	return container.New()
}

// Singleton register a singleton
func Singleton(inst interface{}) {
	instance.Singleton(inst)
}

// Lazy register a singleton lazily
func Lazy(inst interface{}) {
	instance.Lazy(inst)
}

// Factory register a factory
func Factory(inst interface{}) {
	instance.Factory(inst)
}

// Reset remove all bindins
func Reset() {
	instance.Reset()
}

// Make bind receiver
func Make(inst interface{}) {
	instance.Make(inst)
}
