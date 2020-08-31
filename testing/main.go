package main

import (
	gormseeder "github.com/t3hcyberrat/gorm-seeder"
)

// Person ...
type Person struct {
	Latitude           float32 `gormseeder:"lat"`
	Longitude          float32 `gormseeder:"long"`
	CreditCardNumber   string  `gormseeder:"cc_number"`
	CreditCardType     string  `gormseeder:"cc_type"`
	Email              string  `gormseeder:"email"`
	DomainName         string  `gormseeder:"domain_name"`
	IPV4               string  `gormseeder:"ipv4"`
	IPV6               string  `gormseeder:"ipv6"`
	Password           string  `gormseeder:"password"`
	Jwt                string  `gormseeder:"jwt"`
	PhoneNumber        string  `gormseeder:"phone_number"`
	MacAddress         string  `gormseeder:"mac_address"`
	URL                string  `gormseeder:"url"`
	UserName           string  `gormseeder:"username"`
	TitleMale          string  `gormseeder:"title_male"`
	TitleFemale        string  `gormseeder:"title_female"`
	FirstName          string  `gormseeder:"first_name"`
	FirstNameMale      string  `gormseeder:"first_name_male"`
	FirstNameFemale    string  `gormseeder:"first_name_female"`
	LastName           string  `gormseeder:"last_name"`
	Name               string  `gormseeder:"name"`
	Word               string  `gormseeder:"word"`
	Sentence           string  `gormseeder:"sentence"`
	Paragraph          string  `gormseeder:"paragraph"`
	Currency           string  `gormseeder:"currency"`
	AmountWithCurrency string  `gormseeder:"amount_with_currency"`
	UUIDHypenated      string  `gormseeder:"uuid_hyphenated"`
	UUID               string  `gormseeder:"uuid_digit"`
	gormseeder.BaseModel
}

func main() {

	gormseeder.Init("mysql", "root:hacktheplanet@/", "nethkr")
	gormseeder.Run(Person{})

}
