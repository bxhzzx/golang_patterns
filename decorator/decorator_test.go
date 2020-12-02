package decorator

import (
	"log"
	"testing"
)

func TestDecorator(t *testing.T) {
	doubleDecorator := Decorator(Double)
	res := doubleDecorator(222)
	log.Printf("res is %d\n", res)
}
