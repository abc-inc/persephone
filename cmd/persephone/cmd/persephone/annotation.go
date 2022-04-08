package cmd

const (
	Offline = "offline"
)

func Annotate(keys ...string) map[string]string {
	m := make(map[string]string, len(keys))
	for _, k := range keys {
		m[k] = "true"
	}
	return m
}
