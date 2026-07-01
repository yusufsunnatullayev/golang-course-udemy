package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// 1) Not recommended way to initialize a struct
	// p := person{"Yusuf", "Sunnatullayev"}

	// 2) Recommended way to initialize a struct
	// p := person{firstName: "Yusuf", lastName: "Sunnatullayev"}
	// fmt.Println(p)

	// 3) Recommended way to initialize a struct with zero values
	// var p person

	// p.firstName = "Yusuf"
	// p.lastName = "Sunnatullayev"

	// fmt.Println(p)
	// fmt.Printf("%+v", p)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
