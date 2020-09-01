package gormseeder

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker"
	c "github.com/fatih/color"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormwrapper "github.com/zkynet/gorm-wrapper"
	"github.com/zkynetio/logger"
)

const tagName string = "gormseeder"

var dbTag string

// Init ...
func Init(dialect string, connectionString string, DBTag string, sqlLogging bool) error {

	err, _ := logger.Init(&logger.LoggingConfig{
		DefaultLogTag:   "testing-logs",
		DefaultLogLevel: logger.LogLevelInfo,
		WithTrace:       true,
		SimpleTrace:     true,
		PrettyPrint:     true,
		Colors:          true,
		Type:            "stdout",
	})
	if err != nil {
		return err
	}

	dbTag = DBTag
	connectErr := gormwrapper.Connect(dialect, connectionString, dbTag)
	if connectErr != nil {
		connectErr.Log()
		return connectErr
	}

	gormwrapper.ConnectionMap[dbTag].LogMode(sqlLogging)
	if sqlLogging {
		gormwrapper.ConnectionMap[dbTag].SetLogger(&logger.GORMLogger{})
	}

	return nil

}

// Run ...
func Run(Models ...interface{}) {

	for i := range Models {

		err := faker.FakeData(&Models[i])
		if err != nil {
			fmt.Println(err)
		}

		SaveModelToDatabase(Models[i])

		c.Green("Create Type: " + reflect.TypeOf(Models[i]).String())

	}

}

// SaveModelToDatabase ...
func SaveModelToDatabase(M interface{}) error {

	defer func() {

		if r := recover(); r != nil {
			logger.GenericError(logger.GetRecoverError(r)).Log()
		}

	}()

	fmt.Println("Here brah", M)

	err := gormwrapper.Create(dbTag, M)
	if err != nil {
		err.Log()
		return err
	}

	return nil

}
