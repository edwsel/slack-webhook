package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewAction(t *testing.T) {
	action := NewAction("game", "Chess", "button")

	assert.Equal(t, "game", action.Name, "Field name does not match")
	assert.Equal(t, "Chess", action.Text, "Field name does not match")
	assert.Equal(t, "button", action.Type, "Field name does not match")
}

func TestAction_SetName(t *testing.T) {
	action := new(Action)
	action.SetName("game")

	assert.Equal(t, "game", action.Name, "Field name does not match")
}

func TestAction_SetText(t *testing.T) {
	action := new(Action)
	action.SetText("Chess")

	assert.Equal(t, "Chess", action.Text, "Field text does not match")
}

func TestAction_SetStyle(t *testing.T) {
	action := new(Action)
	action.SetStyle("style")

	assert.Equal(t, "style", action.Style, "Field style does not match")
}

func TestAction_SetType(t *testing.T) {
	action := new(Action)
	action.SetType("button")

	assert.Equal(t, "button", action.Type, "Field button does not match")
}

func TestAction_SetValue(t *testing.T) {
	action := new(Action)
	action.SetValue("text value")

	assert.Equal(t, "text value", action.Value, "Field value does not match")
}

func TestAction_SetConfirm(t *testing.T) {
	action := new(Action)
	confirm := new(Confirm)
	action.SetConfirm(*confirm)

	assert.Equal(t, confirm, action.Confirm, "Field confirm does not match")
}

func TestAction_AddOption(t *testing.T) {
	action := new(Action)
	option := *new(Option)
	action.AddOption(option)

	assert.Equal(t, option, action.Options[0], "Field option does not match")
}

func TestAction_AddOptions(t *testing.T) {
	action := new(Action)
	option := *new(Option)
	action.AddOptions([]Option{option})

	assert.Equal(t, []Option{option}, action.Options, "Field option does not match")
}

func TestAction_ToJson(t *testing.T) {
	action := new(Action)
	action.Name = "game"
	action.Text = "Chess"
	action.Type = "button"

	expectedResult := `{"name":"game","text":"Chess","type":"button"}`

	assert.Equal(t, expectedResult, action.ToJson(), "Method's result does not match")
}

