package structAutoGenerate

type IOrmArgs interface {
	SetTable(arg string)
	GetTable() string
	SetFields(arg []string)
	GetFields() []string
	SetWhere(arg [][]interface{})
	GetWhere() [][]interface{}
	SetOrder(arg string)
	GetOrder() string
	SetLimit(arg int)
	GetLimit() int
	SetOffset(arg int)
	GetOffset() int
	SetJoin(arg [][]interface{})
	GetJoin() [][]interface{}
	SetDistinct(arg bool)
	GetDistinct() bool
	SetUnion(arg string)
	GetUnion() string
	SetGroup(arg string)
	GetGroup() string
	SetHaving(arg string)
	GetHaving() string
	SetData(arg interface{})
	GetData() interface{}
}
type OrmArgs struct {
	table    string
	fields   []string
	where    [][]interface{}
	order    string
	limit    int
	offset   int
	join     [][]interface{}
	distinct bool
	union    string
	group    string
	having   string
	data     interface{}
}

func NewOrmArgs() *OrmArgs {
	return new(OrmArgs)
}

func (o *OrmArgs) GetTable() string {
	return o.table
}

func (o *OrmArgs) GetFields() []string {
	return o.fields
}

func (o *OrmArgs) GetWhere() [][]interface{} {
	return o.where
}

func (o *OrmArgs) GetOrder() string {
	return o.order
}

func (o *OrmArgs) GetLimit() int {
	return o.limit
}

func (o *OrmArgs) GetOffset() int {
	return o.offset
}

func (o *OrmArgs) GetJoin() [][]interface{} {
	return o.join
}

func (o *OrmArgs) GetDistinct() bool {
	return o.distinct
}

func (o *OrmArgs) GetUnion() string {
	return o.union
}

func (o *OrmArgs) GetGroup() string {
	return o.group
}

func (o *OrmArgs) GetHaving() string {
	return o.having
}

func (o *OrmArgs) GetData() interface{} {
	return o.data
}
