package main

import (
	"myproject/class"
	"myproject/student"
)

func main() {
	student1 := student.Student{Name: "Alice", Age: 20}
	student2 := student.Student{Name: "Bob", Age: 22}
	student3 := student.Student{Name: "Charlie", Age: 21}
	student4 := student.Student{Name: "David", Age: 23}
	class1 := class.Class{Name: "Class 1", Students: []student.Student{student2, student3}}
	class2 := class.Class{Name: "Class 2", Students: []student.Student{student1, student4}}
	class1.PrintClass()
	class2.PrintClass()
	class1.MoveStudentTo(student2, &class2)
	class1.PrintClass()
	class2.PrintClass()

}
