package crud

import "fmt"

// InvalidAgeError is
type InvalidAgeError struct {
	age int
}

func (i InvalidAgeError) Error() string {
	return fmt.Sprintf("Invalid Age %v, Age can't be <= 0", i.age)
}

// InvalidNameError ..
type InvalidNameError struct {
	name string
}

func (i InvalidNameError) Error() string {
	return fmt.Sprintf("Invalid name %v, Name should contain alphabets only", i.name)
}

// InvalidInputError ..
type InvalidInputError struct {
	InvalidNameError
	InvalidAgeError
}

func (i InvalidInputError) Error() string {
	return i.InvalidNameError.Error() + " " + i.InvalidAgeError.Error()
}

// NotFound ..
type NotFound struct {
	id int
}

func (n NotFound) Error() string {
	return fmt.Sprintf("Student is not found by id %d", n.id)
}
