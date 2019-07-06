package crud

//Student ..
type Student struct {
	id   int
	name string
	age  int
}

var stuMap = make(map[int]Student)

//Create ..
func Create(name string, age int) (int, error) {
	if name == "" {
		return 0, InvalidNameError{name}
	} else if age <= 0 {
		return 0, InvalidAgeError{age}
	}
	id := len(stuMap) + 1
	stu := Student{id, name, age}
	stuMap[id] = stu
	return id, nil
}

// Read ..
func Read(id int) (Student, error) {
	stu, ok := stuMap[id]
	if !ok {
		return stu, NotFound{id}
	}
	return stu, nil
}
