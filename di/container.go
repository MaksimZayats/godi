package di

var globalContainer = NewContainer()

type Container struct {
	// type: T | func(c Container) T | Service[T]
	registeredTypes []any
}

func NewContainer() *Container {
	c := &Container{
		registeredTypes: make([]any, 0),
	}

	return c
}
