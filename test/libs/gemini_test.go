package test

import (
	"dreonbot/libs/gemini"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeminiGenerateContent(t *testing.T) {
	gemini, err := gemini.NewGenAIGemini("AIzaSyDOM0x7DC9ArmETWUt-IkWJQn3nPpihlSY")
	assert.Nil(t, err)

	resp, err := gemini.GenerateContent("Hello")
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}
