// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"github.com/rs/zerolog/log"
)

// Second returns the second argument. It is intended for use in places like
// return, where the second return value of a function call should be returned.
func Second[A any, B any](_ A, b B) B {
	return b
}

// Must panics if the error is non-nil. It is intended for use in variable
// initializations such as
// f := internal.MustNoErr(os.OpenFile("notes.txt", os.O_RDONLY, 0600))
func Must[T any](a T, err error) T {
	MustNoErr(err)
	return a
}

// MustNoErr logs the error message and exits, if the error is non-nil. It is
// intended for use statements that could return an error and continued
// execution is not meaningful.
func MustNoErr(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

// MustOk is similar to Must, but for the "comma ok" idiom. It is intended for
// use in statements like
// str := internal.MustOk(arg.(string))
func MustOk[T any](a T, ok bool) T {
	if !ok {
		log.Fatal().Msgf("Invalid state: %v is not ok", a)
	}
	return a
}

// MustTuple is like Must, but for functions where three values are returned.
func MustTuple[A any, B any](a A, b B, err error) (A, B) {
	MustNoErr(err)
	return a, b
}
