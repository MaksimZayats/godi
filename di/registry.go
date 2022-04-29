package di

func RegisterService[T any](
	service Service[T],
	container ...*Container,
) {
	c := getContainer(container...)
	if service.isSingleton {
		c.registeredTypes = append(c.registeredTypes, service)
	} else if service.fromFactory {
		c.registeredTypes = append(c.registeredTypes, service.factory)
	} else if service.isInstance {
		c.registeredTypes = append(c.registeredTypes, service.instance)
	}
}

func AddScopedByFactory[T any](
	factory func(c *Container) T,
	container ...*Container,
) {
	RegisterService(
		Service[T]{
			ServiceConfig: ServiceConfig{fromFactory: true},
			factory:       factory,
		},
		container...,
	)
}

func AddInstance[T any](
	instance T,
	container ...*Container,
) {
	RegisterService(
		Service[T]{
			ServiceConfig: ServiceConfig{isInstance: true},
			instance:      instance,
		},
		container...,
	)
}

func AddSingletonByFactory[T any](
	factory func(c *Container) T,
	container ...*Container,
) {
	RegisterService(
		Service[T]{
			ServiceConfig: ServiceConfig{
				isSingleton: true,
				fromFactory: true,
			},
			factory: factory,
		},
		container...,
	)

	// c := getContainer(container...)
	// c.registeredTypes = append(
	// 	c.registeredTypes,
	// 	Service[T]{
	// 		ServiceConfig: ServiceConfig{
	// 			isSingleton: true,
	// 			fromFactory: true,
	// 		},
	// 		factory: factory,
	// 	},
	// )
}
