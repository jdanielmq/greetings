
# Saludos en GO

Este paquete proporciana una forma simple de obtener saludos personalizado en Go.

## Installación
Ejecuta el siguiente comando para instalar el paquete:
``` bash
go get -u https://github.com/jdanielmq/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo
``` bash
package main

import (
	"fmt"
	"log"

	"github.com/jdanielmq/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	message, err := greetings.Hello("Juan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
```

este ejemplo importa el paquete github.com/jdanielmq/greetings y se llama a la función Hello para obtener un saludo personalizado, si ocurre un error, se imprime un mensaje de error.

** los errores estan siendo capturado por log de Go.
