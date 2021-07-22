package firsttry_test

import (
	firsttry "firstTry/go-try-1/firstTry"
	"testing"
)

func testTry(t *testing.T) {
	if firsttry.Try() != "Si" {
		t.Fatal("Tu y yo vamos a tener problemas")
	}
}
