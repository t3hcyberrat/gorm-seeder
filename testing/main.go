package main

import (
	"github.com/jinzhu/gorm"
	gormseeder "github.com/t3hcyberrat/gorm-seeder"
)

// Person ...
type Person struct {
	FirstName int
	LastName  string
	gorm.Model
}

func main() {

	gormseeder.Run(Person{})

}
