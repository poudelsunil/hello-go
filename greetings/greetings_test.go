package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T){

	name := "Ram"
	want := regexp.MustCompile(`\b`+name+`\b`)
	message,error := Hello("Ram")
	if( !want.MatchString(message) || error != nil ){
		t.Fatalf(`Hello("Ram") = %q, %v want match for %q, nil`, message, error, want)
	}


}


func TestHelloEmpty(t *testing.T){

	message, error := Hello("")
	if( message != "" || error == nil) {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}

}


func TestHellos(t *testing.T) {

	names := []string{"Ram","Hari","Shyam","Anup","Anuj","Tara","Purna"}
	messages, err := Hellos(names)

	if(len(messages) != 7){
		t.Fatalf("There should be greetings for all");
	}

	if( err != nil){
		t.Fatalf("There should be not be error");
	}
}