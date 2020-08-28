package gormseeder

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	L "github.com/zkynetio/logger"
	"github.com/zkynetio/pi/backend/src/db"
)

func Init(connectionString string, databaseName string) {
	Connect(connectionString, databaseName)
	Connection.LogMode(false)
	Connection.SetLogger(&L.GORMLogger{})
}
func Run(Model interface{}, numberOfRecord int) {
	// Generate rand data based on field?
	generateFields(&Model)

	// Inject field to DB with rand data
	SaveModelToDatabase(&Model)
}

type Test struct {
	ID  string
	Ble string
}

func generateFields(Model interface{}) {

	// TODO ... FILL THE MODEL...
	log.Println("HERE WE DO ALL THE DETECTING AND FILLING")
}
func SaveModelToDatabase(model interface{}) error {
	err := db.Create(&model)
	if err != nil {
		err.Log()
		return err
	}
	return nil
}
