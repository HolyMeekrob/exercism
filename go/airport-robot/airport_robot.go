package airportrobot

import "fmt"

// Greeter defines an entity that can greet people in their own language
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
}

func (language Italian) LanguageName() string {
	return "Italian"
}

func (language Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {
}

func (language Portuguese) LanguageName() string {
	return "Portuguese"
}

func (language Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
