package DI

// Dependency Container
var container map[string]interface{}

// Init container
func init() {
	container = make(map[string]interface{})
}

// Resolving dependency by resource name
func Resolve(name string) interface{} {
	resource := fetchResource(name)
	if resource != nil {
		return resource.(func() interface{})()
	}

	return nil
}

// Injecting singleton resource
func Singleton(name string, resource interface{}) {
	Instance(name, func() interface{} {return resource})
}

// Injecting instance resource
func Instance(name string, factory func() interface{}) {
	container[name] = factory
}

// Injecting resource alias
func Alias(alias string, originName string) {
	resource := fetchResource(originName)
	if resource != nil {
		container[alias] = resource
	}
}

// Fetching resource by name
func fetchResource(name string) interface{} {
	if resource, ok := container[name]; ok {
		return resource
	}

	return nil
}
