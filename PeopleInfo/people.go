package PeopleInfo

import "fmt"

type StructPeople struct {
	fName string
	lName string
	Age   int
}

func InitPeople() []StructPeople {
	person_1 := StructPeople{"Томас", "Шелби", 55}
	person_2 := StructPeople{"Игорь", "Крутой", 69}
}

func ShowPeople(people []StructPeople) {
	for _, person := range people {
		fmt.Println("Имя: %s\nФамилия: %s\nВозраст: %d", person.fName, person.lName, person.Age)
	}
}
