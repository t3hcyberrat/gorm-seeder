package gormseeder

import (
	"fmt"
	"log"
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	L "github.com/zkynetio/logger"
)

// Init ...
func Init(dialect string, connectionString string, databaseName string) {

	Connect(dialect, connectionString, databaseName)
	Connection.LogMode(false)
	Connection.SetLogger(&L.GORMLogger{})

}

// Run ...
func Run(Model interface{}) {

	fmt.Println(Model)

	// Generate rand data
	generateFields(Model)

	// Inject field to DB with rand data
	// SaveModelToDatabase(&Model)

}

func generateFields(Model interface{}) {

	// TODO ... FILL THE MODEL...
	log.Println("HERE WE DO ALL THE DETECTING AND FILLING")

	v := reflect.ValueOf(Model)
	fmt.Println(v)

	typeOfS := v.Type()
	fmt.Println(typeOfS)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s \t \t \tValue: %v \n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

}

// SaveModelToDatabase ...
func SaveModelToDatabase(model interface{}) error {

	err := Create(&model)
	if err != nil {
		err.Log()
		return err
	}

	return nil

}
