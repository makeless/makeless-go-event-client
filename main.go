package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/makeless/makeless-go/event/basic"
	"github.com/makeless/makeless-go/logger/basic"
	"github.com/r3labs/sse/v2"
	"sync"
	"time"
)

func main() {
	var urlFlag = flag.String("url", "", "makeless event source url")
	flag.Parse()

	var colorRed = color.New(color.FgRed, color.Bold)
	var colorGreen = color.New(color.FgGreen, color.Bold)
	var colorWhite = color.New(color.FgWhite)
	var colorHiWhite = color.New(color.FgHiWhite)
	var sep = "   "

	var logger = new(makeless_go_logger_basic.Logger)
	var client = sse.NewClient(*urlFlag)
	client.ReconnectNotify = func(err error, duration time.Duration) {
		if _, err := colorRed.Printf("Error: %s\n", err.Error()); err != nil {
			logger.Fatal(err)
		}
	}

	if err := client.SubscribeRaw(func(event *sse.Event) {
		if event == nil {
			return
		}

		var data = &makeless_go_event_basic.EventData{
			RWMutex: new(sync.RWMutex),
		}

		if err := json.Unmarshal(event.Data, data); err != nil {
			logger.Fatal(fmt.Errorf("event not unmarshallable"))
		}

		if _, err := colorGreen.Printf("%s%s", fillString(fmt.Sprintf("%s:%s", string(event.Event), data.GetId()), 25), sep); err != nil {
			logger.Fatal(err)
		}

		if _, err := colorWhite.Printf("%s%s", time.Now().UTC(), sep); err != nil {
			logger.Fatal(err)
		}

		if _, err := colorHiWhite.Printf("%+v\n", data.GetData()); err != nil {
			logger.Fatal(err)
		}
	}); err != nil {
		logger.Fatal(err)
	}
}

func fillString(content string, length int) string {
	if len(content) >= length {
		return content
	}

	for i := len(content); i <= length; i++ {
		content += " "
	}

	return content
}
