package model

type student struct {
	name   string
	age    int
	fenshu float64
}

func NewStudent(name string, age int, fenshu float64) *student {

	return &student{name, age, fenshu}
}

func GetName(stu student) string {
	return stu.name
}

func GetAge(stu student) int {
	return stu.age
}

func GetFenshu(stu student) float64 {
	return stu.fenshu
}
