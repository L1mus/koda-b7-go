package minitask7

import "fmt"

func CreatePersonData(name, address, phone string) {
	user := NewPerson(name, address, phone)

	fmt.Println("My name is ", user.name)
	fmt.Println("Im From ", user.address)
	fmt.Println("this my contact", user.phone)
	user.SetName("Jhon")
	sayHello := user.Greet()
	fmt.Print(sayHello)
	fmt.Print("My New Name is ", user.name)
}

type Person struct {
	name    string
	address string
	phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		name:    name,
		address: address,
		phone:   phone,
	}
}

func (p *Person) Greet() string {
	return "Hello " + p.name
}

func (p *Person) SetName(newName string) {
	p.name = newName
}
