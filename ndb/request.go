package ndb

const (
	ParamFormat   = "f"
	ParamQuery    = "q"
	ParamTemplate = "t"
)

type Request struct {
	Query    string
	Format   string
	Template string
	Params   map[string]interface{}
}

type Record map[string]interface{}

type Result []Record

type ValueExtractor func(key string) (interface{}, bool)

type RecordExtractor func(keys []string, rse ValueExtractor) Record

type Entity struct {
	Name string
}
