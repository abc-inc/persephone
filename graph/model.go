package graph

import "strings"

type Node struct {
	Count         int64 `json:"count"`
	Relationships map[string]RelProperty
	Type          string   `json:"type"`
	Properties    []string `json:"properties"`
	Labels        []string `json:"labels"`
}

func (n Node) String() string {
	return ":" + strings.Join(n.Labels, ":")
}

type Relationship struct {
	Count      int64                      `json:"count"`
	Type       string                     `json:"type"`
	Properties map[string]NodeRelProperty `json:"properties"`
}

func (r Relationship) String() string {
	return r.Type
}

type RelProperty struct {
	Count      int                        `json:"count"`
	Properties map[string]NodeRelProperty `json:"properties"`
	Direction  string                     `json:"direction"`
	Labels     []string                   `json:"labels"`
}

type NodeProperty struct {
	Existence bool   `json:"existence"`
	Type      string `json:"type"`
	Indexed   bool   `json:"indexed"`
	Unique    bool   `json:"unique"`
}

type NodeRelProperty struct {
	Existence bool   `json:"existence"`
	Type      string `json:"type"`
	Array     bool   `json:"array"`
}