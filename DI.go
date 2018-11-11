package DI

import (
	"reflect"
)

type Container struct {
	// Dependency Container
	container map[string]interface{}
}

var C *Container

// Init container
func init() {
	C = &Container{container: make(map[string]interface{})}
	C.Reset()
}

// Resolving dependency by resource name
func (c *Container) Resolve(name string) interface{} {
	resource := c.fetchResource(name)
	if resource != nil {
		return resource.(func() interface{})()
	}

	return nil
}

// Resolving dependency group by resource names
func (c *Container) ResolveGroup(names []string) []interface{} {
	resources := make([]interface{}, 0, len(names))

	for _, name := range names {
		resources = append(resources, c.Resolve(name))
	}

	return resources
}

// Injecting singleton resource
func (c *Container) Singleton(name string, resource interface{}) {
	c.Instance(name, func() interface{} { return resource })
}

// Injecting instance resource
func (c *Container) Instance(name string, factory func() interface{}) {
	c.container[name] = factory
}

// Injecting resource alias
func (c *Container) Alias(alias string, originName string) {
	resource := c.fetchResource(originName)
	if resource != nil {
		c.container[alias] = resource
	}
}

// Inject singleton resource with tags
func (c *Container) Tag(name string, resource interface{}) {
	// Validate resource type
	reflectType := reflect.TypeOf(resource)
	if reflectType.Kind() != reflect.Ptr {
		return
	}

	reflectTypeElem := reflectType.Elem()

	if reflectTypeElem.Kind() != reflect.Struct {
		return
	}

	// Reference
	reflectValue := reflect.ValueOf(resource).Elem()

	i := 0
	for ; i < reflectTypeElem.NumField(); i++ {
		depName := reflectTypeElem.Field(i).Tag.Get("dep")
		reflectValue.Field(i).Set(reflect.ValueOf(c.Resolve(depName)))
	}

	c.Singleton(name, resource)
}

// Call function with args with dependencies
func (c *Container) Call(method interface{}, argNames []string, argValues []interface{}) []interface{} {
	reflectType := reflect.TypeOf(method)

	if reflectType.Kind() != reflect.Func {
		return nil
	}

	args := make([]reflect.Value, 0, len(argValues))
	for i, argValue := range argValues {
		if argValue == nil {
			args = append(args, reflect.ValueOf(c.Resolve(argNames[i])))
		} else {
			args = append(args, reflect.ValueOf(argValue))
		}
	}

	reflectValue := reflect.ValueOf(method)
	var results []interface{}
	for _, result := range reflectValue.Call(args) {
		results = append(results, result.Interface())
	}

	return results
}

// Reset container
func (c *Container) Reset() {
	c.container = make(map[string]interface{})
}

// Fetching resource by name
func (c *Container) fetchResource(name string) interface{} {
	if resource, ok := c.container[name]; ok {
		return resource
	}

	return nil
}
