package main

import "fmt"

func main() {
	userService1 := InitializeUserService1()
	users, _ := userService1.All()
	fmt.Println(users)

	userService2 := InitializeUserService2()
	users, _ = userService2.All()
	fmt.Println(users)

	userService1Singleton1 := InitializeUserServiceSingleton()
	users, _ = userService1Singleton1.All()
	fmt.Println(users)

	userService1Singleton2 := InitializeUserServiceSingleton()
	users, _ = userService1Singleton2.All()
	fmt.Println(users)
}
