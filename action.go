package slack

const (
	ActionTypeButton = "button"
	ActionTypeSelect = "select"
)

type Action struct {
	Name    string   `json:"name,omitempty"`
	Text    string   `json:"text,omitempty"`
	Style   string   `json:"style,omitempty"`
	Type    string   `json:"type,omitempty"`
	Value   string   `json:"value,omitempty"`
	Confirm *Confirm  `json:"confirm,omitempty"`
	Options []Option `json:"options,omitempty"`
}

func NewAction(name, text, typeAction string) Action {
	return Action{
		Name:name,
		Text:text,
		Type:typeAction,
	}
}

func (action *Action) SetName(name string) *Action {
	action.Name = name
	return action
}

func (action *Action) SetText(text string) *Action {
	action.Text = text
	return action
}

func (action *Action) SetStyle(style string) *Action {
	action.Style = style
	return action
}

func (action *Action) SetType(typeValue string) *Action {
	action.Type = typeValue
	return action
}

func (action *Action) SetValue(value string) *Action {
	action.Value = value
	return action
}

func (action *Action) SetConfirm(confirm Confirm) *Action {
	action.Confirm = &confirm
	return action
}

func (action *Action) AddOption(option Option) *Action {
	action.Options = append(action.Options, option)
	return action
}

func (action *Action) AddOptions(options []Option) *Action {
	action.Options = append(action.Options, options...)
	return action
}

func (action *Action) ToJson() string {
	return toJson(action)
}

