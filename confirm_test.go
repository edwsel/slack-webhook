package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewConfirm(t *testing.T) {
	confirm := NewConfirm("Wouldn't you prefer a good game of chess?")

	assert.Equal(t, "Wouldn't you prefer a good game of chess?",
		confirm.Text, "Field text does not match")
}

func TestConfirm_SetTitle(t *testing.T) {
	confirm := new(Confirm)
	confirm.SetTitle("Are you sure?")

	assert.Equal(t, "Are you sure?",
		confirm.Title, "Field title does not match")
}

func TestConfirm_SetText(t *testing.T) {
	confirm := new(Confirm)
	confirm.SetText("Wouldn't you prefer a good game of chess?")

	assert.Equal(t, "Wouldn't you prefer a good game of chess?",
		confirm.Text, "Field text does not match")
}

func TestConfirm_SetOkText(t *testing.T) {
	confirm := new(Confirm)
	confirm.SetOkText("Yes")

	assert.Equal(t, "Yes",
		confirm.OkText, "Field \"ok\" text does not match")
}

func TestConfirm_SetDismissText(t *testing.T) {
	confirm := new(Confirm)
	confirm.SetDismissText("No")

	assert.Equal(t, "No",
		confirm.DismissText, "Field \"dismiss\" text does not match")
}

func TestConfirm_ToJson(t *testing.T) {
	confirm := new(Confirm)
	confirm.Text = "Wouldn't you prefer a good game of chess?"

	expectedResult := `{"text":"Wouldn't you prefer a good game of chess?"}`

	assert.Equal(t, expectedResult, confirm.ToJson(), "Method's result does not match")
}