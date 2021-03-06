package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("Empty name")
    }

    message := fmt.Sprintf( randomFormat(), name)  //"Hi, %v. Welcome!"
    return message, nil
}

func Hellos(names []string) (map[string]string, error){
    messages := make(map[string]string)

    for _, name := range names {
        message, error := Hello(name)
        if(error != nil){
            return nil, error
        }

        messages[name] = message
    }
    return messages, nil
}

// Go executes init functions automatically at program startup, 
// after global variables have been initialized
func init(){
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string{

    // slice data type => like arry 
    // but its size change dynamically 
    // on inserting & deleting record
    formats := []string{
        "Hi %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}