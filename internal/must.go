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

func First[A any, B any](a A, b B) A {
	return a
}

func Second[A any, B any](a A, b B) B {
	return b
}

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
