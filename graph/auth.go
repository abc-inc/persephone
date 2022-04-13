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

package graph

import (
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type authSchema int

const (
	schemeNone authSchema = iota
	schemeBasic
	schemeKerberos
	schemeBearer
)

func (a authSchema) String() string {
	return []string{"none", "basic", "kerberos", "bearer"}[a]
}

func Auth(auth string) (neo4j.AuthToken, string) {
	typ, cred, _ := strings.Cut(auth, ":")
	switch typ {
	case "": // no credentials at all
		fallthrough
	case schemeNone.String():
		return neo4j.NoAuth(), ""
	case schemeBasic.String():
		u, p, _ := strings.Cut(cred, ":")
		return neo4j.BasicAuth(u, p, ""), u
	case schemeKerberos.String():
		return neo4j.KerberosAuth(cred), ""
	case schemeBearer.String():
		return neo4j.BearerAuth(cred), ""
	default:
		// Assume no scheme is given and fallback to Basic auth.
		return neo4j.BasicAuth(typ, cred, ""), typ
	}
}
