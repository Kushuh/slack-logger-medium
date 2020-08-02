package slackTracker

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	webhook := os.Getenv("WEBHOOK")
	application := os.Getenv("APPLICATION")

	defer func() {
		if err := recover(); err != nil {

		}
	}()

	tracker := Tracker{
		WebHook: webhook,
		Application: application,
	}

	err := tracker.Test("Hello from Go !")
	if err != nil {
		t.Fatalf("Unable to connect to Slack server : %s", err.Error())
	}

	time.Sleep(1 * time.Second)

	err = tracker.Error(fmt.Errorf("Oopsie, the server went woopsie ┐(´д｀)┌ "))
	if err != nil {
		t.Fatalf("Unable to connect to Slack server : %s", err.Error())
	}

	if os.Getenv("BE_CRASHER") == "1" {
		tracker.Fatal(fmt.Errorf("Oopsie, the server went woopsie ┐(´д｀)┌ "))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestTracker")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err = cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
