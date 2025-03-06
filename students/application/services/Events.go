package services

import (
	"apihex01/students/application/repositories"
	"apihex01/students/domain/entities"
)

type Event struct {
	rabbit repositories.IRabbit
}

func NewEvent(rabbit repositories.IRabbit) *Event {
	return &Event{rabbit: rabbit}
}

func (e *Event) Run(student *entities.Student) {
	e.rabbit.SendMessageToBroker(student)
}