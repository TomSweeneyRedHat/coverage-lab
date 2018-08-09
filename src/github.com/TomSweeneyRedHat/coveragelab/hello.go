package coveragelab 

import (
	"fmt"
)

func main() {
	fmt.Println("hello-world")
	CallMePlease()
	CallMePlease()
	CallMePlease()
	CallMePlease()
	msg:=CallMePlease()
	fmt.Println(msg)
}

func CallMePlease() string {
	fmt.Println("Thanks so much for calling!")
	return ("Call back soon!")
}
