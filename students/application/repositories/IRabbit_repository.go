package repositories

import "apihex01/students/domain/entities"

type IRabbit interface {
	SendMessageToBroker(student *entities.Student)
}