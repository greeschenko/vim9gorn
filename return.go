package vim9gorn

type Return struct {
	Value string
}

func NewReturn(value string) *Return {
	return &Return{Value: value}
}

func (r *Return) Generate() string {
	return "return " + r.Value
}
