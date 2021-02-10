package main

import "fmt"

func main() {

	var x = map[string]string{
		"Name":    "Viral",
		"Surname": "Shastri",
	}

	fmt.Println(x)
	fmt.Println(x["Name"])
	fmt.Println(x["Surname"])

	v, ok := x["Middlename"]
	fmt.Printf("`%s`\n", v)
	fmt.Println(ok)

	if v, ok := x["Middlename"]; ok {
		fmt.Printf("`%s`\n", v)
		fmt.Println(ok)
	}

	for key, value := range x {
		fmt.Printf("%s = %s\n", key, value)
	}

}
