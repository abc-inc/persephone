package internal

import (
	"fmt"
	"log"
)

func MustNoErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func MustOk[T any](a T, ok bool) T {
	if !ok {
		panic(fmt.Sprintf("Invalid state: %v is not ok\n", a))
	}
	return a
}

func Must[T any](a T, err error) T {
	MustNoErr(err)
	return a
}

func MustTuple[A any, B any](a A, b B, err error) (A, B) {
	MustNoErr(err)
	return a, b
}
