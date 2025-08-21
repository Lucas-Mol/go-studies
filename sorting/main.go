package main

import (
	"fmt"
	"sort"
)

// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func main() {
	// ==== Ints and Strings
	// numbers := []int{5, 3, 4, 1, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)

	// ==== Function and Interface

	// people := []Person{
	// 	{"Alice", 30},
	// 	{"Bob", 25},
	// 	{"Ana", 35},
	// }

	// fmt.Println("Unsorted by age:", people)
	// sort.Sort(ByAge(people))
	// fmt.Println("Sorted by age:", people)

	// ==== Generic Function and Interface

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Ana", 35},
	}

	fmt.Println("Unsorted:", people)

	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(ageAsc).Sort(people)

	fmt.Println("Sorted by age (ascending):", people)

	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	By(ageDesc).Sort(people)

	fmt.Println("Sorted by age (descending):", people)

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)

	fmt.Println("Sorted by name:", people)

	// ==== Slice

	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)
}
