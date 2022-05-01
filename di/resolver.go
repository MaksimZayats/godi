package di

import (
	"errors"
)

func Get[T any]() (T, error) {
	return GetFromContainer[T](globalContainer)
}

func GetFromContainer[T any](c *Container) (T, error) {
	for i := 0; i < len(c.registeredTypes); i++ {
		switch typedValue := c.registeredTypes[i].(type) {
		case T:
			return typedValue, nil
		case func(c *Container) T:
			return typedValue(c), nil
		case Service[T]:
			if typedValue.isSingleton {
				if typedValue.fromFactory {
					instance := typedValue.factory(c)
					c.registeredTypes[i] = instance
					return instance, nil
				}
			}
		}
	}
	return *new(T), errors.New("unregistered type")
}

func MustGet[T any]() T {
	return MustGetFromContainer[T](globalContainer)
}

func MustGetFromContainer[T any](c *Container) T {
	value, err := GetFromContainer[T](c)
	if err != nil {
		panic(err)
	}
	return value
}
