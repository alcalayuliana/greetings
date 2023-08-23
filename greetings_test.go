package greetings

import (
	"testing"
	"regexp"
)

func TestHello(t *testing.T){
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)//Busca una coincidencia  exacta con el nombre con esta expresión regular
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err!= nil{
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q,nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil{
        t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
    }
}