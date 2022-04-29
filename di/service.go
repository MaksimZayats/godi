package di

type ServiceConfig struct {
	isSingleton bool
	isInstance  bool
	fromFactory bool
}

type Service[T any] struct {
	ServiceConfig
	factory  func(c *Container) T
	instance T
}
