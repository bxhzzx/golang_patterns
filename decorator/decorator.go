package decorator

import "log"

type Object func(int) int

func Double(a int) int {
	return a * 2
}

func Decorator(object Object) Object {
	return func(a int) int {
		log.Println("before decorator")
		res := object(a)
		log.Println("after decorator")
		return res
	}
}
