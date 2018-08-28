package slack

type Message struct {
	Text        string       `json:"text,omitempty"`
	Username    string       `json:"username,omitempty"`
	Mrkdwn      bool         `json:"mrkdwn,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

func NewMessage(text string) Message {
	return Message{Text: text}
}

func (message *Message) SetText(text string) *Message {
	message.Text = text
	return message
}

func (message *Message) SetUsername(username string) *Message {
	message.Username = username
	return message
}

func (message *Message) UseMrkdwn() *Message {
	message.Mrkdwn = true
	return message
}

func (message *Message) AddAttachment(attachment Attachment) *Message {
	message.Attachments = append(message.Attachments, attachment)
	return message
}

func (message *Message) AddAttachments(attachments []Attachment) *Message {
	message.Attachments = append(message.Attachments, attachments...)
	return message
}

func (message *Message) ToJson() string {
	return toJson(message)
}
