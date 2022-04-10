package internal

import (
	"github.com/rs/zerolog/log"
)

func MustNoErr(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func MustOk[T any](a T, ok bool) T {
	if !ok {
		log.Fatal().Msgf("Invalid state: %v is not ok", a)
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
