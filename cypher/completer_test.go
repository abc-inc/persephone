package cypher

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func Test_findVars(t *testing.T) {
	Equal(t, []string{}, findVars("MATCH () RETURN 1"))
	Equal(t, []string{}, findVars("MATCH (:A) RETURN 1"))
	Equal(t, []string{"n"}, findVars("MATCH (n) RETURN n"))
	Equal(t, []string{"a"}, findVars("MATCH (a:A) RETURN a"))
	Equal(t, []string{"a", "b"}, findVars("MATCH (a)--(b) RETURN b"))

	Equal(t, []string{}, findVars("MATCH ()-[:R]-() RETURN 1"))
	Equal(t, []string{"r"}, findVars("MATCH ()-[r]-() RETURN r"))
	Equal(t, []string{"r"}, findVars("MATCH ()-[r:R]-() RETURN r"))

	// Equal(t, []string{"n", "c"}, findVars("MATCH (n) RETURN count(n) AS c, count(n) AS n"))
	Equal(t, []string{"a", "r1", "b", "r2", "e"},
		findVars("MATCH (a:A)<-[r1:Rel]-(b)-[r2]-(:C)-[:r3]-()--(e:E {name: 'Name'}) RETURN a"))
}
