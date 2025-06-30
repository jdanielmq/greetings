package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("El nombre viene vacio")
	}
	message := fmt.Sprintf(randomFormnat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}

func randomFormnat() string {
	formats := []string{
		"Hola, %v! Bienvenido! ",
		"hola soy , %v! Bienvenido ha esta aventura! ",
		"Hola, %v! ",
	}
	return formats[rand.Intn(len(formats))]
}
