package di

func GetContainer(container ...*Container) *Container {
	switch len(container) {
	case 0:
		return globalContainer
	case 1:
		return container[0]
	default:
		panic("Unsupported number of arguments")
	}
}
