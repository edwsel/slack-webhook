package slack

import (
	"os"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"testing"
	"github.com/stretchr/testify/assert"
)

const configFileName = "test-config.yaml"

type testConfig struct {
	ApiKey string `yaml:"api_key"`
	Chanel string `yaml:"chanel"`
}

func TestNew(t *testing.T) {
	config := getTestConfig()

	slack := New(config.ApiKey)

	assert.Equal(t, config.ApiKey, slack.apiKey, "Field api key does not match")
}

func TestSlack_UseDebug(t *testing.T) {
	slack := new(Slack)
	slack.UseDebug()

	assert.Equal(t, true, slack.debug, "Field debug does not match")
}

func TestSlack_Send(t *testing.T) {
	config := getTestConfig()

	message := *new(Message)
	message.Text = "Slack auto test message"

	slack := new(Slack)
	slack.apiKey = config.ApiKey
	slack.debug = true
	slack.logger = new(slackLogger)

	err := slack.Send(config.Chanel, message)

	assert.Nil(t, err, "When you send an error has occurred")
}

func getTestConfig() testConfig {
	testConfig := new(testConfig)

	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		return *testConfig
	}

	yamlData, err := ioutil.ReadFile(configFileName)

	if err != nil {
		return *testConfig
	}

	yaml.Unmarshal(yamlData, testConfig)

	return *testConfig
}