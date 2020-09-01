package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormseeder "github.com/t3hcyberrat/gorm-seeder"
)

// Person ...
type Person struct {
	gorm.Model
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	Jwt                string  `faker:"jwt"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`
}

// XXX
type XXX struct {
	Name string
	gorm.Model
}

func main() {

	// db, err := gorm.Open("mysql", "root:hacktheplanet@/nethkr")
	// if err != nil {
	// 	panic(err)
	// }

	// db.SingularTable(true)

	// // Migrate the schema
	// db.AutoMigrate(&Person{})
	// db.AutoMigrate(&XXX{})

	err := gormseeder.Init("mysql", "", "seeder-connection", true)
	if err != nil {
		os.Exit(1)
	}

	gormseeder.Run(
		Person{},
		XXX{},
	)

}
