package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandsEvents(commandsChannel <-chan *slacker.CommandEvent) {
	for event := range commandsChannel {
		fmt.Println("Commands Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5210512052882-5210743210755-aKO869BK5VBzH0nzgq3FgXRQ")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A056908AJPN-5223156184625-6b54c693fdc44540ae8c78aba6b40389a4c6002b3cee524e72bda292440bc3c9")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandsEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
