package lang

import (
	"regexp"
	"strings"
)

func EscapeCypher(str string) string {
	prefix := ""
	if strings.HasPrefix(str, ":") {
		prefix = ":"
	}

	content := strings.TrimPrefix(str, prefix)
	if ok, err := regexp.MatchString(`^[A-Za-z][A-Za-z0-9_]*$`, content); err == nil && ok {
		return str
	} else {
		return prefix + "`" + strings.ReplaceAll(content, "`", "``") + "`"
	}
}
