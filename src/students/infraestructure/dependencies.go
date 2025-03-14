package infraestructure

import "apihex01/src/students/infraestructure/adapters"

var (
	mysql  *MySQL
	rabbit *adapters.Rabbit
)

func InitDepedencies() {
	mysql = NewMySQL()
	rabbit = adapters.NewRabbitMq()
}

func GetMySQL() *MySQL {
	return mysql
}

func GetRabbit() *adapters.Rabbit {
	return rabbit
}
