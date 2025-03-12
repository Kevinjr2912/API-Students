package infraestructure

var (
	mysql *MySQL
	rabbit *Rabbit
)

func InitDepedencies() {
	mysql = NewMySQL()
	rabbit = NewRabbitMq()
}

func GetMySQL() *MySQL{
	return mysql
}

func GetRabbit() *Rabbit {
	return rabbit
}

