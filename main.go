package slack

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"errors"
)

const headerApplicationJson = "application/json"
const slackHostBase = "https://hooks.slack.com/services/%s"

type Slack struct {
	apiKey string
	debug  bool
	logger Logger
}

func New(apiKey string) Slack {
	return Slack{
		apiKey: apiKey,
		debug:  false,
		logger: new(slackLogger),
	}
}

func (slack *Slack) SetLogger(logger Logger) {
	slack.logger = logger
}

func (slack *Slack) log(url, requestData, responseData, error string) {
	logString := fmt.Sprintf(`
Url: %s
Request data: %s
Respounse data: %s
Error %s
`, url, requestData, responseData, error)

	slack.logger.Log(logString)
}

func (slack *Slack) Send(to string, message Message) error {
	url := fmt.Sprintf(slackHostBase, to)
	preparedMessage := message.ToJson()

	return slack.request(url, preparedMessage)
}

func (slack *Slack) UseDebug() {
	slack.debug = true
}

func (slack *Slack) request(url, data string) error {
	preparedData := []byte(data)
	dataReader := bytes.NewReader(preparedData)
	response, err := http.Post(url, headerApplicationJson, dataReader)
	defer response.Body.Close()

	if err != nil {
		if slack.debug {
			slack.log(url, data, "", err.Error())
		}

		return err
	}

	responseDataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		if slack.debug {
			slack.log(url, data, "", err.Error())
		}

		return err
	}

	responseData := string(responseDataByte)

	if response.StatusCode != 200 {
		errHttp := errors.New("http error")
		if slack.debug {
			slack.log(url, data, responseData, errHttp.Error())
		}

		return errHttp
	}

	if slack.debug {
		slack.log(url, data, responseData, "")
	}

	return nil
}
