package main

import (
	log "github.com/thinkboy/log4go"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Student struct {
	Person

	School string
	Loan   float32
}

type Employee struct {
	Person

	Company string
	Money   float32
}

type iStudent interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

func (m *Student) test() {
	log.Debug("(%v, %v, %v, %v, %v)", m.Name, m.Age, m.Address, m.School, m.Loan)

}

func (m *Employee) test() {
	log.Debug("(%v, %v, %v, %v, %v)", m.Name, m.Age, m.Address, m.Company, m.Money)
}
