package main

import (
    "fmt"
    "github.com/hello-go/greetings"
)

func main() {
    message := greetings.Hello("Ben")
    fmt.Println(message)
}
