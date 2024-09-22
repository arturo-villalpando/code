package classes

import "strings"

type Woman struct {
	Person
}

func NewWoman(name string, lastname string, age int32) *Woman {
	return &Woman{
		Person{name, lastname, age},
	}
}

/**
 * Example of Polymorphism
 */
func (p Woman) Scream() string {
	return strings.ToUpper(p.Name) + " " + strings.ToUpper(p.Lastname)
}
