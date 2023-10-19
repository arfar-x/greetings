package greetings

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// varDump will print out any number of variables given to it
// e.g. varDump("test", 1234)
func varDump(myVar ...interface{}) {
	fmt.Printf("%v\n", myVar)
}

// dd will print out variables given to it (like varDump()) but
// will also stop execution from continuing.
func dd(myVar ...interface{}) {
	varDump(myVar...)
	os.Exit(1)
}
