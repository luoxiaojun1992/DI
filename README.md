# DI
Dependency Injection in Go

# Functions
* Singleton Register
* Instance Register
* Alias Register
* Tag Register
* Resolve Single Dependency
* Resolve Grouped Dependency
* Resolve Tag Dependency

# Test Result
![](docs/test_result.png)

# Examples
```go
var container = DI.C

//Singleton Register
container.Singleton("UserService", struct {
		name string
	}{name: "hello"})
container.Resolve("UserService")

//Instance Register
container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
container.Resolve("UserService")
	
//Alias Register
container.Alias("UserServ", "UserService")
container.Resolve("UserServ")

//Tag Register
container.Tag("TagDemo", &struct {
		Name interface{} `dep:"UserService"`
	}{Name: "test"})
container.Resolve("TagDemo")

//Resolve Grouped Dependency
container.Singleton("UserService", struct {
		name string
	}{name: "new user"})
container.Singleton("GoodsService", struct {
		name string
	}{name: "new goods"})
container.Singleton("OrderService", struct {
		name string
	}{name: "new order"})
container.ResolveGroup([]string{"UserService", "GoodsService", "OrderService"})
```
