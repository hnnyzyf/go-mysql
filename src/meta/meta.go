package meta

//获得指定域的信息
type Meta interface {
	//获得Meta所包含的列的列表
	GetMetaType() string
}

//存储表的Column信息
type ColumnTable struct {
	Field map[string]int
	Type  []SqlType
	Null  []bool

	size int
}

func NewColumnTable() *ColumnTable {
	return &ColumnTable{
		Field: make(map[string]int),
		Type:  []SqlType{},
		Null:  []bool{},
		size:  0,
	}
}

var IsYes0No = map[string]bool{
	"YES": true,
	"NO":  false,
}

func (c *ColumnTable) GetMetaType() string {
	return "column"
}

//实现GetField
func (c *ColumnTable) GetField() []string {
	var res []string
	for key, _ := range c.Field {
		res = append(res, key)
	}
	return res
}

//实现GetType
func (c *ColumnTable) GetType(ob string) SqlType {
	idx := c.Field[ob]
	return c.Type[idx]
}

//实现IsNull
func (c *ColumnTable) IsNull(ob string) bool {
	idx := c.Field[ob]
	return c.Null[idx]
}

//实现IsNull
func (c *ColumnTable) HasColumn(ob string) bool {

	if _, ok := c.Field[ob]; ok {
		return true
	} else {
		return false
	}
}

func (c *ColumnTable) Insert(f string, ct string, b string) {
	c.Field[f] = c.size
	c.Type = append(c.Type, getsqltype(ct))
	c.Null = append(c.Null, IsYes0No[b])
	c.size++
}
