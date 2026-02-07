/*
	Встраивание структур

	Дана структура Human (с произвольным набором полей и методов).

	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

	Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

package main

import "fmt"

type Human struct {
	surname string
	name    string
	age     int
}

func NewHuman(surname, name string, age int) *Human {
	return &Human{
		surname: surname,
		name:    name,
		age:     age,
	}
}

func (h *Human) GetSurname() string {
	return h.surname
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

type Action struct {
	Human
	city string
}

func NewAction(h Human, city string) *Action {
	return &Action{
		Human: h,
		city:  city,
	}
}

func (a *Action) GetCity() string {
	return a.city
}

func main() {
	human := NewHuman("Ivanov", "Ivan", 100)
	action := NewAction(*human, "Spb")

	fmt.Printf("Surname: %v,\nName: %v,\nAge: %v,\nCity: %v.\n", action.GetSurname(), action.GetName(), action.age, action.GetCity())
}
