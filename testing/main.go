package main

import (
	"github.com/jinzhu/gorm"
	gormseeder "github.com/t3hcyberrat/gorm-seeder"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

func main() {

	p1 := Person{
		FirstName: "Peter",
		LastName:  "Gore",
	}

	gormseeder.Run(p1)

}
