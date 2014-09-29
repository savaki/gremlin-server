package types

import "fmt"

type Map map[string]interface{}

func (m Map) Href() string {
	if v, ok := m["href"]; ok {
		return fmt.Sprintf("%v", v)
	} else {
		return ""
	}
}

func (m Map) Url() string {
	return m.Href()
}

type Url interface {
	Url() string
}
