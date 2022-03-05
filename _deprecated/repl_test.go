package _deprecated

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func Test_findVars(t *testing.T) {
	Equal(t, []string{}, FindVars("MATCH () RETURN 1"))
	Equal(t, []string{}, FindVars("MATCH (:A) RETURN 1"))
	Equal(t, []string{"n"}, FindVars("MATCH (n) RETURN n"))
	Equal(t, []string{"a"}, FindVars("MATCH (a:A) RETURN a"))
	Equal(t, []string{"a", "b"}, FindVars("MATCH (a)--(b) RETURN b"))

	Equal(t, []string{}, FindVars("MATCH ()-[:R]-() RETURN 1"))
	Equal(t, []string{"r"}, FindVars("MATCH ()-[r]-() RETURN r"))
	Equal(t, []string{"r"}, FindVars("MATCH ()-[r:R]-() RETURN r"))

	Equal(t, []string{"c", "n"}, FindVars("MATCH (n) RETURN count(n) AS c, count(n) AS n"))
	Equal(t, []string{"a", "b", "e", "r1", "r2"},
		FindVars("MATCH (a:A)<-[r1:Rel]-(b)-[r2]-(:C)-[:r3]-()--(e:E {name: 'Name'}) RETURN a"))
}
