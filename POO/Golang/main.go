package main

import (
	"fmt"

	p "poo/classes"
)

func main() {
	// Person
	person := p.NewPerson("Arturo", "Villalpando", 3)
	// Print name
	fmt.Println(person.GetFullname())
	// Print age
	fmt.Println("My age is: ", person.GetAge())
	// Is the person birthday
	person.IsPersonBirthday()
	// Print the new age after birthday
	fmt.Println("Age after birthday is: ", person.GetAge())
	// Scream
	fmt.Println(person.Scream())
	// Space
	fmt.Println()
	// Woman
	woman := p.NewWoman("Gabriela", "Villalpando", 4)
	// Print name
	fmt.Println(woman.GetFullname())
	// Print age
	fmt.Println("My age is: ", woman.GetAge())
	// Is the person birthday
	person.IsPersonBirthday()
	// Print the new age after birthday
	fmt.Println("Age after birthday is: ", woman.GetAge())
	// Scream
	fmt.Println(woman.Scream())
}
