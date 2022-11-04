package main

import "fmt"

type Person struct {
	Phone   string
	Name    string
	Surname string
	Gender  int
}

type Option func(*Person)

func Phone(phone string) Option {
	return func(p *Person) {
		p.Phone = phone
	}
}

func NewPerson(opts ...Option) *Person {
	p := &Person{
		Phone: "00000000000000",
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func main() {
	person := NewPerson(Phone("12345678"))
	fmt.Println(person.Phone)
}
