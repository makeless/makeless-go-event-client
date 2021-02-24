package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	makeless_go_event_basic "github.com/makeless/makeless-go/event/basic"
	"github.com/makeless/makeless-go/logger/basic"
	"github.com/r3labs/sse/v2"
	"sync"
	"time"
)

func main() {
	var colorGreen = color.New(color.FgGreen, color.Bold)
	var colorWhite = color.New(color.FgWhite)
	var colorHiWhite = color.New(color.FgHiWhite)

	var logger = new(makeless_go_logger_basic.Logger)
	var client = sse.NewClient("https://localhost:3003/api/auth/company/1/group/1/message/event")

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

		if _, err := colorGreen.Printf("%s:%s\t\t", string(event.Event), data.GetId()); err != nil {
			logger.Fatal(err)
		}

		if _, err := colorHiWhite.Printf("%s\t\t", data.GetData()); err != nil {
			logger.Fatal(err)
		}

		if _, err := colorWhite.Printf("%s\n", time.Now().UTC().String()); err != nil {
			logger.Fatal(err)
		}
	}); err != nil {
		logger.Fatal(err)
	}
}
