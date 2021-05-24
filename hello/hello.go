package main

import (
	"fmt"
	"log"

    "github.com/hello-go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	
	message, err := greetings.Hello("Ben")
	if(err != nil){
		log.Fatal(err)
	}
	
	fmt.Println(message)


	names := []string{"Ram","Hari","Shyam","Anup","Anuj","Tara","Purna"}

	messages, err := greetings.Hellos(names)
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(messages)
	
}
