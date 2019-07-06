package main

import (
	"fmt"
	"go-crud-app/crud"
	"log"
)

func main() {

	//reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=================================")
		fmt.Println("Please select action")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Exit")
		fmt.Println("=================================")
		var action int
		fmt.Scanf("%d", &action)
		switch action {
		case 1:
			var name string
			var age int
			fmt.Print("Enter Name: ")
			fmt.Scanf("%s", &name)
			fmt.Print("Enter age: ")
			fmt.Scanf("%d", &age)
			id, err := crud.Create(name, age)
			if ageErr, ok := err.(crud.InvalidAgeError); ok {
				log.Println(ageErr)
				fmt.Print("Please enter valid Age: ")
				fmt.Scanf("%d", age)
				crud.Create(name, age)
			} else if nameErr, ok := err.(crud.InvalidNameError); ok {
				log.Println(nameErr)
				fmt.Print("Please enter valid Name : ")
				fmt.Scanf("%s", name)
				crud.Create(name, age)
			}
			fmt.Println("Student Created with id: ", id)
		case 2:
			var id int
			fmt.Print("Enter Id: ")
			fmt.Scanf("%d", &id)
			stu, err := crud.Read(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Found student", stu)
		case 3:
		case 4:
		case 5:

		}

	}
}
