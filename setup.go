package slackWatcher

import (
	"fmt"
	"github.com/Alvarios/poster"
	"log"
	"os"
	"runtime/debug"
	"time"
)

type Tracker struct {
	WebHook string
	Application string
}

func (t *Tracker) Test(message string) error {
	_, err := poster.Post(t.WebHook, map[string]interface{}{"text" : message})
	return err
}

// Avoid code repetition : let's handle every common operation in this unexported function.
func (t *Tracker) send(color, message string) error {
	env := os.Getenv("ENV")
	// If no ENV is specified, assume we are in development mode, so we don't want to flood Slack uselessly.
	if env == "" {
		return nil
	}

	_, perr := poster.Post(
		t.WebHook,
		map[string]interface{}{
			"text" : fmt.Sprintf("%s - %s", t.Application, env),
			"attachments": []map[string]interface{}{
				{
					"color": color,
					"text": fmt.Sprintf(
						"*Message*\n%s\n\n*Stack*\n```%s```\n\n*Time*\n%s",
						message,
						string(debug.Stack()),
						time.Now().Format("2006-01-02 03:04:05"),
					),
				},
			},
		},
	)

	// An unexpected error happened when sending our message to Slack.
	return perr
}

func (t *Tracker) Error(err error) error {
	return t.send("#ff9300", err.Error())
}

func (t *Tracker) Fatal(err error) {
	_ = t.send("#ff3232", err.Error())
	log.Fatal(err.Error())
}