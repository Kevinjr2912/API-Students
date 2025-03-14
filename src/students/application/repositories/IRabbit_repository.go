package repositories

import "apihex01/src/students/domain/entities"

type IRabbit interface {
	SendMessageToBroker(student *entities.StudentCredentials)
}