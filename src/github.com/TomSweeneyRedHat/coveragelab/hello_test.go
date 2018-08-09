package coveragelab

import (
	"fmt"
	"testing"
)

func TestApiFail(t *testing.T) {
	msg := CallMePlease()
	fmt.Printf("Failure test msg: [%s]\n", msg)
	if msg == "Not Real Error" {
		t.Errorf("We called the wrong api! [%s]", msg)
	}
}

func TestApiSucceed(t *testing.T) {
	msg := CallMePlease()
	fmt.Printf("Success test msg: [%s]\n", msg)
	if msg != "Call back soon!" {
		t.Errorf("We called the wrong api! [%s]", msg)
	}
}
