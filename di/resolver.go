package di

func Get[T any]() T {
	return GetFromContainer[T](globalContainer)
}

func GetFromContainer[T any](c *Container) T {
	for i := 0; i < len(c.registeredTypes); i++ {
		switch typedValue := c.registeredTypes[i].(type) {
		case T:
			return typedValue
		case func(c *Container) T:
			return typedValue(c)
		case Service[T]:
			if typedValue.isSingleton {
				if typedValue.fromFactory {
					instance := typedValue.factory(c)
					c.registeredTypes[i] = instance
					return instance
				}
			}
		}
	}

	panic("Unregistered type")
}
