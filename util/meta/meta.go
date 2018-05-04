package meta

import "strings"

type Attribute struct {
	Type   SqlType
	IsNull bool
}

func NewAttribute(t string, n string) *Attribute {
	return &Attribute{
		Type:   getSqlType(t),
		IsNull: isYes0No[n],
	}
}

func getSqlType(t string) SqlType {
	s := strings.Split(t, "(")
	return SqlDict[s[0]]
}

var isYes0No = map[string]bool{
	"YES": true,
	"NO":  false,
}

type Meta struct {
	meta  []*Attribute
	alias []string
	size  int
}

func NewMeta() *Meta {
	return &Meta{
		meta:  []*Attribute{},
		alias: []string{},
		size:  0,
	}
}

func (m *Meta) Add(f string, t string, n string) {
	m.meta = append(m.meta, NewAttribute(t, n))
	m.alias = append(m.alias, f)
	m.size++
}

func (m *Meta) AddMeta(other *Meta) {
	for i := 0; i < other.size; i++ {
		m.meta = append(m.meta, other.meta[i])
		m.alias = append(m.alias, other.alias[i])
		m.size++
	}
}

func (m *Meta) GetField() []string {
	return m.alias
}

func (m *Meta) GetSqlType(column string) SqlType {
	for idx, col := range m.alias {
		if strings.ToLower(col) == strings.ToLower(column) {
			return m.meta[idx].Type
		}
	}

	panic("无法确定column的类型")
	return Type_unknown
}
