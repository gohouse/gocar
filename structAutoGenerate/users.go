package structAutoGenerate

type IUsers interface {
	SetFields(arg []interface{})
	GetFields() []interface{}
	SetResult(arg []map[string]interface{})
	GetResult() []map[string]interface{}
	SetTagName(arg string)
	GetTagName() string
	SetTagIgnoreName(arg string)
	GetTagIgnoreName() string
}
type Users struct {
	Fields        []interface{}
	Result        []map[string]interface{}
	TagName       string
	TagIgnoreName string
}

func NewUsers() *Users {
	return new(Users)
}

func (o *Users) SetFields(arg []interface{}) {
	o.Fields = arg
}

func (o *Users) GetFields() []interface{} {
	return o.Fields
}

func (o *Users) SetResult(arg []map[string]interface{}) {
	o.Result = arg
}

func (o *Users) GetResult() []map[string]interface{} {
	return o.Result
}

func (o *Users) SetTagName(arg string) {
	o.TagName = arg
}

func (o *Users) GetTagName() string {
	return o.TagName
}

func (o *Users) SetTagIgnoreName(arg string) {
	o.TagIgnoreName = arg
}

func (o *Users) GetTagIgnoreName() string {
	return o.TagIgnoreName
}
