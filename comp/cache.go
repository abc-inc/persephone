package comp

import (
	"github.com/abc-inc/merovingian/comp/rule"
	"github.com/abc-inc/merovingian/types"
)

type Cache struct {
	Map map[types.Type][]Item
}

func (c Cache) CalculateItems(ri rule.Info, query string) []Item {
	return nil
}

func (c Cache) Complete(ris []rule.Info, query string) (its []Item) {
	for _, ri := range ris {
		if cached, ok := c.Map[ri.Type]; ok {
			its = append(its, cached...)
			continue
		}
		its = append(its, c.CalculateItems(ri, query)...)
	}
	return
}
