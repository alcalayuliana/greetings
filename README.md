# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personales

## Instalación
Ejecuta el siguiente comando para instalar el paquete:
``` bash
go get -u github.com/alcalayuliana/greetings
```



## Uso
Aquí tiene un ejemplo de cómo utilizar el paquete en tu código:


``` go
package main
import (
    "fmt"
    "github.com/alcalayuliana/greetings"
)

func main() {
    message, err := greetings.Hello("Alex")
    if err != nil {
        fmt.Println("Ocurrión un error:", err)
        return
    }
    fmt.Println("message")
}

```

Este ejemplo importa el paquete github.com/alcalayuliana/greetings  y llama a la función Hello y devuelve un saludo personalizado.
Si ocurre un error, se imprime un mensaje de error.