package slack

type Option struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewOption(text, value string) Option {
	return Option{
		Text: text,
		Value: value,
	}
}

func (option *Option) SetText(text string) *Option {
	option.Text = text
	return option
}

func (option *Option) SetValue(value string) *Option {
	option.Value = value
	return option
}

func (option *Option) ToJson() string {
	return toJson(option)
}
