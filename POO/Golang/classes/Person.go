package classes

import (
	"strings"
)

/**
 * Structs is like a class
 */
type Person struct {
	Name     string
	Lastname string
	Age      int32
}

/**
 * Construct
 */
func NewPerson(name string, lastname string, age int32) *Person {
	return &Person{
		Name:     name,
		Lastname: lastname,
		Age:      age,
	}
}

/**
 * Getters
 */
func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetLastname() string {
	return p.Lastname
}

func (p *Person) GetFullname() string {
	return p.GetName() + " " + p.GetLastname()
}

func (p *Person) GetAge() int32 {
	return p.Age
}

/**
 * Actions
 */
func (p *Person) IsPersonBirthday() {
	p.Age += 1
}

/**
 * Example of Polymorphism
 */
func (p Person) Scream() string {
	return strings.ToUpper(p.Name)
}
