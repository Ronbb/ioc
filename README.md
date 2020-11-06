[![GoDoc](https://godoc.org/github.com/ronbb/ioc?status.svg)](https://godoc.org/github.com/ronbb/ioc)
[![Go Report Card](https://goreportcard.com/badge/github.com/ronbb/ioc)](https://goreportcard.com/report/github.com/ronbb/ioc)

# ioc

A lightweight ioc container for golang.

## Install

```bash
go get -u github.com/ronbb/ioc
```

## Usage

### Singleton

Register a singleton.

```go
ioc.Singleton(&instance)
// or
ioc.Singleton(func() Abstraction {
    return Implementation
})
```

It's better to use a pointer for unneccessary copy.

### Lazy

Register a lazy singleton.

```go
ioc.Lazy(func() Abstraction {
    return Implementation
})
```

### Factory

Register a factory.

```go
ioc.Factory(func() Abstraction {
    return Implementation
})
```

### Reset

Reset the default container.

```go
ioc.Reset()
```

### NewContainer

Create a new container.

```go
container := ioc.NewContainer()
// container.Singleton(...)
// container.Factory(...)
// container.Lazy(...)
// container.Reset()
// container.Make(...)
```

## Addtional

This project is insiped by [github.com/golobby/container](https://github.com/golobby/container).
