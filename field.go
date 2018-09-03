package slack

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

func NewField(title string) *Field {
	return &Field{
		Title: title,
	}
}

func (field *Field) SetTitle(title string) *Field {
	field.Title = title
	return field
}

func (field *Field) SetValue(value string) *Field {
	field.Value = value
	return field
}

func (field *Field) UseShort() *Field {
	field.Short = true
	return field
}

func (field *Field) ToJson() string {
	return toJson(field)
}