package gormseeder

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	L "github.com/zkynetio/logger"
)

type ModelWrapper struct{}

// Init ...
func Init(dialect string, connectionString string, databaseName string) {

	Connect(dialect, connectionString, databaseName)
	Connection.LogMode(false)
	Connection.SetLogger(&L.GORMLogger{})

}

// Run ...
func Run(Model interface{}) {

	// Generate rand data
	generateFields(Model)

	// Inject field to DB with rand data
	// SaveModelToDatabase(&Model)

}

func generateFields(Model interface{}) {

	// TODO ... FILL THE MODEL...
	composedStruct := CreateObject(Model)

	for i, v := range composedStruct {
		fmt.Println(i, v)
		switch v.(type) {
		case string:
			composedStruct[i] = faker.Email()
		}
	}

	fmt.Println(composedStruct)

}

// CreateObject ...
func CreateObject(Model interface{}) []interface{} {

	modelValues := reflect.ValueOf(Model)
	ret := make([]interface{}, modelValues.NumField())

	for i := 0; i < modelValues.NumField(); i++ {
		ret[i] = modelValues.Field(i).Interface()
	}

	return ret

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
