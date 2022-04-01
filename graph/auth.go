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
		return neo4j.NoAuth(), ""
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
