package slack

import "log"

type Logger interface {
	Log(args ...interface{})
}

type slackLogger struct {
}

func (logger *slackLogger) Log(args ...interface{})  {
	log.Printf("%+v\n", args)
}
