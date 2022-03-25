package internal

import "log"

func MustNoErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Must[T any](a T, err error) T {
	MustNoErr(err)
	return a
}
