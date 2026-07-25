package class

import "myproject/student"

type Class struct {
	Name string

	Students []student.Student
}

func (c *Class) AddStudent(s student.Student) {
	c.Students = append(c.Students, s)
}

func (c *Class) MoveStudentTo(s student.Student, target ClassManage) {
	for i, student := range c.Students {
		if student == s {
			c.Students = append(c.Students[:i], c.Students[i+1:]...)
			target.AddStudent(s)
			break
		}
	}
}

func (c *Class) PrintClass() {
	println("Class Name:", c.Name)
	println("Students:")
	for _, student := range c.Students {
		println(" - ", student.GetName(), " (", student.GetAge(), ")")
	}
}

type ClassManage interface {
	AddStudent(s student.Student)
	MoveStudentTo(s student.Student, c ClassManage)
	PrintClass()
}
