package main

import (
	"fmt"
	"sort"
)

// User Structure
type User struct {
	name  string
	age   int
	marks []int
}

// ByName implements sort.Interface for []Person based on the Age field.
type ByName []User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }

func main() {

	users := []User{
		{"Viral", 22, []int{65, 21, 45}},
		{"Shastri", 105, []int{80, 70, 45}},
		{"Manish", 51, []int{55, 45, 47}},
	}

	fmt.Println("Not Sorted:\n", users)

	for _, user := range users {
		sort.Ints(user.marks)
	}

	fmt.Println("After Sorting Each Users Marks:")
	fmt.Println(users)

	sort.Sort(ByName(users))

	fmt.Println("Sorted by Name:\n", users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})

	fmt.Println("Sorted by Age:\n", users)
}
