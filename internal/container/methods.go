package container

// Singleton register a singleton
func (c *container) Singleton(resolver interface{}) {
	c.bind(resolver, singleton)
}

// Lazy register a singleton lazily
func (c *container) Lazy(resolver interface{}) {
	c.bind(resolver, lazy)
}

// Factory register a factory
func (c *container) Factory(resolver interface{}) {
	c.bind(resolver, factory)
}

// Reset remove all bindins
func (c *container) Reset() {
	for k := range c.bindings {
		delete(c.bindings, k)
	}
}

// Make bind receiver
func (c *container) Make(receiver interface{}) {
	c.make(receiver)
}
