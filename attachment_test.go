package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachment(t *testing.T) {
	attachment := NewAttachment("Would you like to play a game?")

	assert.Equal(t, "Would you like to play a game?", attachment.Title,
		"Field text does not match")
}

func TestAttachment_SetCallbackId(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetCallbackId("game_selection")

	assert.Equal(t, "game_selection", attachment.CallbackID,
		"Field callback id does not match")
}

func TestAttachment_SetAttachmentType(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetAttachmentType("default")

	assert.Equal(t, "default", attachment.AttachmentType,
		"Field attachment type does not match")
}

func TestAttachment_SetFallback(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetFallback("If you could read this message, you'd be choosing something fun to do right now.")

	assert.Equal(t, "If you could read this message, you'd be choosing something fun to do right now.",
		attachment.Fallback,
		"Field callback does not match")
}

func TestAttachment_SetColor(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetColor("#3AA3E3")

	assert.Equal(t, "#3AA3E3", attachment.Color,
		"Field color does not match")
}

func TestAttachment_SetPretext(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetPretext("New ticket from Andrea Lee")

	assert.Equal(t, "New ticket from Andrea Lee", attachment.Pretext,
		"Field pretext does not match")
}

func TestAttachment_SetAuthorName(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetAuthorName("Bobby Tables")

	assert.Equal(t, "Bobby Tables", attachment.AuthorName,
		"Field author name does not match")
}

func TestAttachment_SetAuthorLink(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetAuthorLink("http://flickr.com/bobby/")

	assert.Equal(t, "http://flickr.com/bobby/", attachment.AuthorLink,
		"Field author link does not match")
}

func TestAttachment_SetAuthorIcon(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetAuthorIcon("http://flickr.com/icons/bobby.jpg")

	assert.Equal(t, "http://flickr.com/icons/bobby.jpg", attachment.AuthorIcon,
		"Field author icon does not match")
}

func TestAttachment_SetTitle(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetTitle("Slack API Documentation")

	assert.Equal(t, "Slack API Documentation", attachment.Title,
		"Field title does not match")
}

func TestAttachment_SetTitleLink(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetTitleLink("https://api.slack.com/")

	assert.Equal(t, "https://api.slack.com/", attachment.TitleLink,
		"Field title link does not match")
}

func TestAttachment_SetText(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetText("Optional text that appears within the attachment")

	assert.Equal(t, "Optional text that appears within the attachment", attachment.Text,
		"Field callback id does not match")
}

func TestAttachment_SetImageURL(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetImageURL("http://my-website.com/path/to/image.jpg")

	assert.Equal(t, "http://my-website.com/path/to/image.jpg", attachment.ImageURL,
		"Field image url does not match")
}

func TestAttachment_SetThumbURL(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetThumbURL("http://example.com/path/to/thumb.png")

	assert.Equal(t, "http://example.com/path/to/thumb.png", attachment.ThumbURL,
		"Field thumb url does not match")
}

func TestAttachment_SetFooter(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetFooter("Slack API")

	assert.Equal(t, "Slack API", attachment.Footer,
		"Field footer does not match")
}

func TestAttachment_SetFooterIcon(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetFooterIcon("https://platform.slack-edge.com/img/default_application_icon.png")

	assert.Equal(t, "https://platform.slack-edge.com/img/default_application_icon.png", attachment.FooterIcon,
		"Field footer icon does not match")
}

func TestAttachment_SetTs(t *testing.T) {
	attachment := new(Attachment)
	attachment.SetTs(123456789)

	assert.Equal(t, 123456789, attachment.Ts,
		"Field ts does not match")
}

func TestAttachment_AddFiled(t *testing.T) {
	attachment := new(Attachment)
	attachment.AddFiled(*new(Field))

	assert.Equal(t, *new(Field), attachment.Fields[0],
		"Field fields does not match")
}

func TestAttachment_AddFields(t *testing.T) {
	attachment := new(Attachment)
	attachment.AddFields([]Field{*new(Field)})

	assert.Equal(t, []Field{*new(Field)}, attachment.Fields,
		"Field fields does not match")
}

func TestAttachment_ToJson(t *testing.T) {
	attachment := new(Attachment)
	attachment.Text = "Would you like to play a game?"

	expectedResult := `{"text":"Would you like to play a game?"}`

	assert.Equal(t, expectedResult, attachment.ToJson(),
		"Method's result does not match")
}