package event

import "reflect"

type FormatEvent struct {
	Format string
	Sep    string
}

type Subscriber[E any] func(e E)

var subByType = make(map[string][]Subscriber[any])

func Subscribe[E any](e E, s Subscriber[E]) {
	subByType[typeOf(e)] = append(subByType[typeOf(e)], func(e interface{}) {
		s(e.(E))
	})
}

func Publish[E any](e E) {
	for _, s := range subByType[typeOf(e)] {
		s(e)
	}
}

func typeOf(e interface{}) string {
	return reflect.TypeOf(e).Name()
}
