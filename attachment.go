package slack

type Attachment struct {
	CallbackID     string   `json:"callback_id,omitempty"`
	AttachmentType string   `json:"attachment_type,omitempty"`
	Fallback       string   `json:"fallback,omitempty"`
	Color          string   `json:"color,omitempty"`
	Pretext        string   `json:"pretext,omitempty"`
	AuthorName     string   `json:"author_name,omitempty"`
	AuthorLink     string   `json:"author_link,omitempty"`
	AuthorIcon     string   `json:"author_icon,omitempty"`
	Title          string   `json:"title,omitempty"`
	TitleLink      string   `json:"title_link,omitempty"`
	Text           string   `json:"text,omitempty"`
	Fields         []Field  `json:"fields,omitempty"`
	Actions        []Action `json:"actions,omitempty"`
	ImageURL       string   `json:"image_url,omitempty"`
	ThumbURL       string   `json:"thumb_url,omitempty"`
	Footer         string   `json:"footer,omitempty"`
	FooterIcon     string   `json:"footer_icon,omitempty"`
	Ts             int      `json:"ts,omitempty"`
}

func NewAttachment(title string) Attachment {
	return Attachment{
		Title: title,
	}
}

func (attachment *Attachment) SetCallbackId(callbackID string) *Attachment {
	attachment.CallbackID = callbackID
	return attachment
}

func (attachment *Attachment) SetAttachmentType(attachmentType string) *Attachment {
	attachment.AttachmentType = attachmentType
	return attachment
}

func (attachment *Attachment) SetFallback(fallback string) *Attachment {
	attachment.Fallback = fallback
	return attachment
}

func (attachment *Attachment) SetColor(color string) *Attachment {
	attachment.Color = color
	return attachment
}

func (attachment *Attachment) SetPretext(pretext string) *Attachment {
	attachment.Pretext = pretext
	return attachment
}

func (attachment *Attachment) SetAuthorName(authorName string) *Attachment {
	attachment.AuthorName = authorName
	return attachment
}

func (attachment *Attachment) SetAuthorLink(authorLink string) *Attachment {
	attachment.AuthorLink = authorLink
	return attachment
}

func (attachment *Attachment) SetAuthorIcon(authorIcon string) *Attachment {
	attachment.AuthorIcon = authorIcon
	return attachment
}

func (attachment *Attachment) SetTitle(title string) *Attachment {
	attachment.Title = title
	return attachment
}

func (attachment *Attachment) SetTitleLink(titleLink string) *Attachment {
	attachment.TitleLink = titleLink
	return attachment
}

func (attachment *Attachment) SetText(text string) *Attachment {
	attachment.Text = text
	return attachment
}

func (attachment *Attachment) SetImageURL(imageURL string) *Attachment {
	attachment.ImageURL = imageURL
	return attachment
}

func (attachment *Attachment) SetThumbURL(thumbURL string) *Attachment {
	attachment.ThumbURL = thumbURL
	return attachment
}

func (attachment *Attachment) SetFooter(footer string) *Attachment {
	attachment.Footer = footer
	return attachment
}

func (attachment *Attachment) SetFooterIcon(footerIcon string) *Attachment {
	attachment.FooterIcon = footerIcon
	return attachment
}

func (attachment *Attachment) SetTs(ts int) *Attachment {
	attachment.Ts = ts
	return attachment
}

func (attachment *Attachment) AddFiled(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, field)
	return attachment
}

func (attachment *Attachment) AddFields(fields []Field) *Attachment {
	attachment.Fields = append(attachment.Fields, fields...)
	return attachment
}

func (attachment *Attachment) ToJson() string {
	return toJson(attachment)
}