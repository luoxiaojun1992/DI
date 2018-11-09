package tests

import (
	"github.com/luoxiaojun1992/DI"
	"log"
	"testing"
)

var container = DI.C

func Test_SingletonResolve(t *testing.T) {
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

func Benchmark_SingletonResolve(b *testing.B) {
	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserService").(struct{ name string }).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_InstanceResolve(b *testing.B) {
	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserService").(struct{ name string }).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasSingletonResolve(b *testing.B) {
	container.Singleton("UserService", struct {
		name string
	}{name: "hello"})
	container.Alias("UserServ", "UserService")

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserServ").(struct{ name string }).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasInstanceResolve(b *testing.B) {
	container.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
	container.Alias("UserServ", "UserService")

	i := 0
	for ; i <= b.N; i++ {
		if container.Resolve("UserServ").(struct{ name string }).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_ResolveGroup(b *testing.B) {
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
			log.Fatal("User name is incorrect")
		}

		if services[1].(struct{ name string }).name != "new goods" {
			log.Fatal("Goods name is incorrect")
		}

		if services[2].(struct{ name string }).name != "new order" {
			log.Fatal("Order name is incorrect")
		}
	}
}

func Benchmark_TagResolve(b *testing.B) {
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
			log.Fatal("User name is incorrect")
		}
	}
}
