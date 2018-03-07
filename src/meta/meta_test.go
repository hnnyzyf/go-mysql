package meta

import "testing"

func Test_Table(t *testing.T) {
	s := NewSimpleTable()
	s.Insert("a", 2, true)
	t.Log(s)

	t.Log(s.Column())
	t.Log(s.GetType("s"))
	t.Log(s.IsNull("a"))
}
