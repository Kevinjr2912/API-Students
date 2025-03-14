package services

import (
	"apihex01/src/students/application/repositories"
	"apihex01/src/students/domain/entities"
)

type Event struct {
	rabbit repositories.IRabbit
}

func NewEvent(rabbit repositories.IRabbit) *Event {
	return &Event{rabbit: rabbit}
}

func (e *Event) Run(student *entities.StudentCredentials) {
	e.rabbit.SendMessageToBroker(student)
}