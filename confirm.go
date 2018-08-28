package slack

type Confirm struct {
	Title       string `json:"title,omitempty"`
	Text        string `json:"text,omitempty"`
	OkText      string `json:"ok_text,omitempty"`
	DismissText string `json:"dismiss_text,omitempty"`
}

func NewConfirm(text string) Confirm {
	return Confirm{
		Text: text,
	}
}

func (confirm *Confirm) SetTitle(title string) *Confirm {
	confirm.Title = title
	return confirm
}

func (confirm *Confirm) SetText(text string) *Confirm {
	confirm.Text = text
	return confirm
}

func (confirm *Confirm) SetOkText(okText string) *Confirm {
	confirm.OkText = okText
	return confirm
}

func (confirm *Confirm) SetDismissText(dismissText string) *Confirm {
	confirm.DismissText = dismissText
	return confirm
}

func (confirm *Confirm) ToJson() string {
	return toJson(confirm)
}
