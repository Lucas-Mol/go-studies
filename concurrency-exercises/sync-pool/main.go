package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Creating a new person")
			return &person{}
		},
	}

	// getting an object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18
	fmt.Println("Got Person:", person1)

	pool.Put(person1)
	fmt.Println("Returned person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got another Person:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got another Person:", person3)

	person3.name = "Doe"
	person3.age = 20

	// returning objects to the pool again
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Returned people to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got another Person:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got another Person:", person5)
}
