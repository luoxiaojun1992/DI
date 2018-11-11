package tests

import (
	"github.com/luoxiaojun1992/DI"
	"testing"
)

var container = DI.C

func Test_SingletonResolve(t *testing.T) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	userService := container.Resolve("UserService")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_InstanceResolve(t *testing.T) {
	container.Reset()

	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})

	userService := container.Resolve("UserService")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_AliasSingletonResolve(t *testing.T) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})
	container.Alias("UserServ", "UserService")

	userService := container.Resolve("UserServ")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_AliasInstanceResolve(t *testing.T) {
	container.Reset()

	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
	container.Alias("UserServ", "UserService")

	userService := container.Resolve("UserServ")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_ResolveGroup(t *testing.T) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "new user"})

	container.Singleton("GoodsService", struct {
		name string
	}{name: "new goods"})

	container.Singleton("OrderService", struct {
		name string
	}{name: "new order"})

	services := container.ResolveGroup([]string{"UserService", "GoodsService", "OrderService"})

	if services[0] == nil {
		t.Fatal("UserService not found")
	}

	user, ok := services[0].(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}

	if user.name != "new user" {
		t.Fatal("User name is incorrect")
	}

	if services[1] == nil {
		t.Fatal("GoodsService not found")
	}

	goods, ok := services[1].(struct{ name string })
	if !ok {
		t.Fatal("GoodsService type error")
	}

	if goods.name != "new goods" {
		t.Fatal("Goods name is incorrect")
	}

	if services[2] == nil {
		t.Fatal("OrderService not found")
	}

	order, ok := services[2].(struct{ name string })
	if !ok {
		t.Fatal("OrderService type error")
	}

	if order.name != "new order" {
		t.Fatal("Order name is incorrect")
	}
}

func Test_TagResolve(t *testing.T) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	container.Tag("TagDemo", &struct {
		Name interface{} `dep:"UserService"`
	}{Name: "test"})

	tagDemo := container.Resolve("TagDemo")
	if tagDemo == nil {
		t.Fatal("TagDemo not found")
	}

	val, ok := tagDemo.(*struct {
		Name interface{} `dep:"UserService"`
	})
	if !ok {
		t.Fatal("TagDemo type error")
	}

	userService, ok := val.Name.(struct{ name string })
	if !ok {
		t.Fatal("UserService type error")
	}
	if userService.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_Call(t *testing.T)  {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	method := func(userService struct {name string}) string {return userService.name}
	result := container.Call(method, []string{"UserService"}, []interface{}{nil})

	if result[0] != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_CallSpec(t *testing.T)  {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	method := func(args ...interface{}) interface{} {return args[0].(struct{name string}).name}

	if container.CallSpec(method, []string{"UserService"}, []interface{}{nil}) != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Benchmark_SingletonResolve(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserService").(struct{ name string }).name != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_InstanceResolve(b *testing.B) {
	container.Reset()

	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserService").(struct{ name string }).name != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasSingletonResolve(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})
	container.Alias("UserServ", "UserService")

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserServ").(struct{ name string }).name != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasInstanceResolve(b *testing.B) {
	container.Reset()

	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
	container.Alias("UserServ", "UserService")

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserServ").(struct{ name string }).name != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_ResolveGroup(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "new user"})

	container.Singleton("GoodsService", struct {
		name string
	}{name: "new goods"})

	container.Singleton("OrderService", struct {
		name string
	}{name: "new order"})

	i := 0
	for ; i <= b.N; i++ {
		services := container.ResolveGroup([]string{"UserService", "GoodsService", "OrderService"})

		if services[0].(struct{ name string }).name != "new user" {
			b.Fatal("User name is incorrect")
		}

		if services[1].(struct{ name string }).name != "new goods" {
			b.Fatal("Goods name is incorrect")
		}

		if services[2].(struct{ name string }).name != "new order" {
			b.Fatal("Order name is incorrect")
		}
	}
}

func Benchmark_TagResolve(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	container.Tag("TagDemo", &struct {
		Name interface{} `dep:"UserService"`
	}{Name: "test"})

	i := 0
	for ; i <= b.N; i++ {
		tagDemo := container.Resolve("TagDemo")
		if tagDemo.(*struct {
			Name interface{} `dep:"UserService"`
		}).Name.(struct{ name string }).name != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_Call(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	method := func(userService struct{ name string }) string { return userService.name }

	i := 0
	for ; i <= b.N; i++ {
		result := container.Call(method, []string{"UserService"}, []interface{}{nil})

		if result[0] != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_CallSpec(b *testing.B) {
	container.Reset()

	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	method := func(args ...interface{}) interface{} { return args[0].(struct{name string}).name }

	i := 0
	for ; i <= b.N; i++ {
		if container.CallSpec(method, []string{"UserService"}, []interface{}{nil}) != "hello" {
			b.Fatal("User name is incorrect")
		}
	}
}
