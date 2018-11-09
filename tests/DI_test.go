package tests

import (
	"testing"
	"github.com/luoxiaojun1992/DI"
	"log"
)

func Test_SingletonResolve(t *testing.T) {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	userService := DI.Resolve("UserService")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{name string})
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_InstanceResolve(t *testing.T) {
	DI.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})

	userService := DI.Resolve("UserService")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{name string})
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_AliasSingletonResolve(t *testing.T) {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})
	DI.Alias("UserServ", "UserService")

	userService := DI.Resolve("UserServ")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{name string})
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Test_AliasInstanceResolve(t *testing.T) {
	DI.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
	DI.Alias("UserServ", "UserService")

	userService := DI.Resolve("UserServ")
	if userService == nil {
		t.Fatal("UserService not found")
	}

	val, ok := userService.(struct{name string})
	if !ok {
		t.Fatal("UserService type error")
	}

	if val.name != "hello" {
		t.Fatal("User name is incorrect")
	}
}

func Benchmark_SingletonResolve(b *testing.B)  {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})

	i := 0
	for ;i <= b.N;i++ {
		if DI.Resolve("UserService").(struct{name string}).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_InstanceResolve(b *testing.B)  {
	DI.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})

	i := 0
	for ;i <= b.N;i++ {
		if DI.Resolve("UserService").(struct{name string}).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasSingletonResolve(b *testing.B)  {
	DI.Singleton("UserService", struct {
		name string
	}{name: "hello"})
	DI.Alias("UserServ", "UserService")

	i := 0
	for ;i <= b.N;i++ {
		if DI.Resolve("UserServ").(struct{name string}).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}

func Benchmark_AliasInstanceResolve(b *testing.B)  {
	DI.Instance("UserService", func() interface{} {
		return struct {
			name string
		}{name: "hello"}
	})
	DI.Alias("UserServ", "UserService")

	i := 0
	for ;i <= b.N;i++ {
		if DI.Resolve("UserServ").(struct{name string}).name != "hello" {
			log.Fatal("User name is incorrect")
		}
	}
}
