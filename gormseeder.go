package gormseeder

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	L "github.com/zkynetio/logger"
)

const tagName string = "gormseeder"

// Init ...
func Init(dialect string, connectionString string, databaseName string) {

	Connect(dialect, connectionString, databaseName)
	Connection.LogMode(false)
	Connection.SetLogger(&L.GORMLogger{})

}

// Run ...
func Run(Model interface{}) {

	// Generate rand data
	modelMap := generateFields(Model)

	// Inject field to DB with rand data
	SaveModelToDatabase(Model, modelMap)

}

func generateFields(Model interface{}) map[string]interface{} {

	modelMap := CreateObject(Model)

	for k := range modelMap {

		switch k {
		case "lat":
			modelMap[k] = faker.Latitude()
		case "long":
			modelMap[k] = faker.Longitude()
		case "cc_number":
			modelMap[k] = faker.CCNumber()
		case "cc_type":
			modelMap[k] = faker.CCType()
		case "email":
			modelMap[k] = faker.Email()
		case "domain_name":
			modelMap[k] = faker.DomainName()
		case "ipv4":
			modelMap[k] = faker.IPv4()
		case "ipv6":
			modelMap[k] = faker.IPv6()
		case "password":
			modelMap[k] = faker.Password()
		case "jwt":
			modelMap[k] = faker.Jwt()
		case "phone_number":
			modelMap[k] = faker.Phonenumber()
		case "mac_address":
			modelMap[k] = faker.MacAddress()
		case "url":
			modelMap[k] = faker.URL()
		case "username":
			modelMap[k] = faker.Username()
		case "title_male":
			modelMap[k] = faker.TitleMale()
		case "title_female":
			modelMap[k] = faker.TitleFemale()
		case "first_name":
			modelMap[k] = faker.FirstName()
		case "first_name_male":
			modelMap[k] = faker.FirstNameMale()
		case "first_name_female":
			modelMap[k] = faker.FirstNameFemale()
		case "last_name":
			modelMap[k] = faker.LastName()
		case "name":
			modelMap[k] = faker.Name()
		case "word":
			modelMap[k] = faker.Word()
		case "sentence":
			modelMap[k] = faker.Sentence()
		case "paragraph":
			modelMap[k] = faker.Paragraph()
		case "currency":
			modelMap[k] = faker.Currency()
		case "amount_with_currency":
			modelMap[k] = faker.AmountWithCurrency()
		case "uuid_hyphenated":
			modelMap[k] = faker.UUIDHyphenated()
		case "uuid_digit":
			modelMap[k] = faker.UUIDDigit()
		}

	}

	return modelMap

}

// CreateObject ...
func CreateObject(Model interface{}) map[string]interface{} {

	modelValues := reflect.ValueOf(Model)
	result := make(map[string]interface{})

	for i := 0; i < modelValues.NumField(); i++ {

		tag := modelValues.Type().Field(i).Tag.Get(tagName)

		if tag == "" {
			tag = modelValues.Type().Field(i).Name
		}

		result[tag] = modelValues.Field(i).Interface()

	}

	return result

}

// SaveModelToDatabase ...
func SaveModelToDatabase(M interface{}, modelMap map[string]interface{}) error {

	x := reflect.TypeOf(&M)
	// y := reflect.Indirect(&x)

	fmt.Println(x)
	// fmt.Println(y)

	// err := CreateFromMap(&x, modelMap)
	// if err != nil {
	// 	err.Log()
	// 	return err
	// }

	return nil

}
