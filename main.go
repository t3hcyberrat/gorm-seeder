package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// Get model
	// Detect fields types in model
	modelDetect := detect()
	fmt.Println(modelDetect)

	// Generate rand data based on field?
	generate := generateFields()
	fmt.Println(generate)

	// Inject field to DB with rand data
	injectFields := inject()
	fmt.Println(injectFields)

}

func setConnection(db *gorm.DB) {

	fmt.Println(db)

}

func detect() string {

	return "detected"

}

func generateFields() string {

	return "Generating fields"

}
func inject() string {

	return "Injecting fields into DB"

}
