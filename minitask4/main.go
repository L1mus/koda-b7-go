package main

import "fmt"

type Education struct {
	name  string
	major string
}

func main() {
	type Identity struct {
		name      string
		picture   string
		email     string
		age       int
		phone     string
		isMarried bool
		education []Education
	}
	user := Identity{
		name:      "Ali",
		picture:   "https://i.pinimg.com/control1/736x/a5/78/03/a57803e6f37f73382bf89b007a4b5954.jpg",
		email:     "limustadji@gmail.com",
		age:       27,
		phone:     "+6285156534946",
		isMarried: false,
		education: []Education{
			{name: "Universitas Pakuan", major: "S1 Ilmu Komputer"},
		},
	}

	fmt.Println(user)
}
