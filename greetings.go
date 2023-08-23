package greetings	

import (
	"fmt"
	"errors"
	"math/rand"
) 

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	

	if name == "" {
		return "", errors.New("El nombre no puede estar vacío")
	}
	//Devuelve un saludo que contiene incluyendo el nombre en un mensaje
    message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) //mapa clave de tip string y valor de tipo string

	for _, name := range names {
		message, err := Hello(name) 
		
        if err!= nil {
            return nil, err
        }

        messages[name] = message //mapa clave: nombre, valor: mensaje
	}
	return messages, nil
}

//randomFormat devuelve uno de varios formatos de mensajes de saludo
//se selecciona de forma aleatoria

func randomFormat()string {
	format := []string{
	"Hola, %v! Bienvenido!", 
	"¡Que bueno verte, %v!",
    "Saludo, %v! Encantado de conocerte!",
}
 return format[rand.Intn(len(format))]
}