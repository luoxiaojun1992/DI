# DI
Dependency Injection in Go

# Functions
* Singleton Register
* Instance Register
* Alias Register
* Resolve Single Dependency
* Resolve Grouped Dependency

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
