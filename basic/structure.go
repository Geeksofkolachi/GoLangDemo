package main

type student struct {
	name      string
	className string
	age       int
}

func main() {
	student1 := student{name: "Zain", className: "A", age: 24}
	student2 := student{"Haseeb", "A", 24}
	println(student1.name)
	println(student2.name)
}
