package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	message := NewMessage("Would you like to play a game?")

	assert.Equal(t, "Would you like to play a game?", message.Text,
		"Field text does not match")
}

func TestMessage_SetText(t *testing.T) {
	message := new(Message)
	message.SetText("Would you like to play a game?")

	assert.Equal(t, "Would you like to play a game?", message.Text,
		"Field text does not match")
}

func TestMessage_SetUsername(t *testing.T) {
	message := new(Message)
	message.SetUsername("Belson Bot")

	assert.Equal(t, "Belson Bot", message.Username,
		"Field username does not match")
}

func TestMessage_UseMrkdwn(t *testing.T) {
	message := new(Message)
	message.UseMrkdwn()

	assert.Equal(t, true, message.Mrkdwn,
		"Field mrkdwn does not match")
}

func TestMessage_AddAttachment(t *testing.T) {
	message := new(Message)
	message.AddAttachment(*new(Attachment))

	assert.Equal(t, *new(Attachment), message.Attachments[0],
		"Field attachments does not match")
}

func TestMessage_AddAttachments(t *testing.T) {
	message := new(Message)
	message.AddAttachments([]Attachment{*new(Attachment)})

	assert.Equal(t, []Attachment{*new(Attachment)}, message.Attachments,
		"Field attachments does not match")
}

func TestMessage_ToJson(t *testing.T) {
	message := new(Message)
	message.Text = "Would you like to play a game?"

	expectedResult := `{"text":"Would you like to play a game?"}`

	assert.Equal(t, expectedResult, message.ToJson(), "Method's result does not match")
}