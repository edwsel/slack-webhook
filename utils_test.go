package slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToJson(t *testing.T)  {
	data := map[string]interface{}{
		"test": "data",
	}
	expectedResult := `{"test":"data"}`
	json := toJson(data)

	assert.Equal(t, expectedResult, json, "Method's result does not match")
}
