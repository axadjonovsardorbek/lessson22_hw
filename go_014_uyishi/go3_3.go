package main

import (
	"fmt"
	"time"
)

type User struct {
	Id     int
	Name   string
	Events []Event
}

type Event struct {
	Id int
	EventName string
	Time      time.Time
}

func AddEvent(u *User, e []Event){
	u.Events = e
}

func RemoveEvent(e *[]Event, id int){
	for i := 0; i < len(*e); i++{
		if id == (*e)[i].Id{
			e = deleteElement(e, i)
		}
	}
}

func main() {


	u := []User{
		{
			Id:   1,
			Name: "Mohir",
		},
	}

	e := []Event{
		{
			Id: 1,
			EventName: "Hamkorlar bilan uchrashuv",
			Time: time.Date(2024,4,20, 15, 20, 0, 0, time.UTC),
		},
		{
			Id: 2,
			EventName: "Qarindoshlar bilan uchrashuv",
			Time: time.Date(2024,4,22, 15, 20, 0, 0, time.UTC),
		},
		{
			Id: 3,
			EventName: "Muhim dars",
			Time: time.Date(2024,5,20, 15, 20, 0, 0, time.UTC),
		},
	}
	

	AddEvent(&u[0], e)
	RemoveEvent(&e, 2)

	fmt.Println(u)

}
