package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	field := NewField("Priority")

	assert.Equal(t, "Priority", field.Title, "Field title does not match")
}

func TestField_SetTitle(t *testing.T) {
	field := new(Field)
	field.SetTitle("Priority")

	assert.Equal(t, "Priority", field.Title, "Field title does not match")
}

func TestField_SetValue(t *testing.T) {
	field := new(Field)
	field.SetValue("High")

	assert.Equal(t, "High", field.Value, "Field value does not match")
}

func TestField_UseShort(t *testing.T) {
	field := new(Field)
	field.UseShort()

	assert.Equal(t, true, field.Short, "Field short does not match")
}

func TestField_ToJson(t *testing.T) {
	field := new(Field)
	field.Title = "Priority"

	expectedResult := `{"title":"Priority"}`

	assert.Equal(t, expectedResult, field.ToJson(), "Method's result does not match")
}
