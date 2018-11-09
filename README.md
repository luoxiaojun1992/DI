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
//Singleton Register
DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})
DI.Resolve("UserService")

//Instance Register
DI.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
DI.Resolve("UserService")
	
//Alias Register
DI.Alias("UserServ", "UserService")
DI.Resolve("UserServ")

//Tag Register
DI.Tag("TagDemo", &struct {
		Name interface{} `dep:"UserService"`
	}{Name: "test"})
DI.Resolve("TagDemo")

//Resolve Grouped Dependency
DI.Singleton("UserService", struct {
		name string
	}{name: "new user"})
DI.Singleton("GoodsService", struct {
		name string
	}{name: "new goods"})
DI.Singleton("OrderService", struct {
		name string
	}{name: "new order"})
DI.ResolveGroup([]string{"UserService", "GoodsService", "OrderService"})
```
