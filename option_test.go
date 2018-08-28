package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewOption(t *testing.T) {
	option := NewOption("Hearts", "hearts")

	assert.Equal(t, "Hearts", option.Text, "Field text does not match")
	assert.Equal(t, "hearts", option.Value, "Field value does not match")
}

func TestOption_SetText(t *testing.T) {
	option := new(Option)
	option.SetText("Hearts")

	assert.Equal(t, "Hearts",
		option.Text, "Field text does not match")
}

func TestOption_SetValue(t *testing.T) {
	option := new(Option)
	option.SetValue("hearts")

	assert.Equal(t, "hearts",
		option.Value, "Field value does not match")
}

func TestOption_ToJson(t *testing.T) {
	option := new(Option)
	option.Text = "Hearts"
	option.Value = "hearts"

	expectedResult := `{"text":"Hearts","value":"hearts"}`

	assert.Equal(t, expectedResult, option.ToJson(), "Method's result does not match")

}