package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-6319028162357-6324424866868-b8l43QvJapq12hiLSwIeSEZD")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A069FR1QC59-6324415959860-bce1cf3fd3dafb60e808b632889062a6142332cade5a0b822127ab326374f4ae")
	bot:= slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("Toi sinh nam <year>", &slacker.CommandDefinition{
		Description: "May tinh tuoi",
		Examples: []string{"Toi sinh nam 2000",
		"Toi sinh nam 2003",
	},

		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year:= request.Param("year")
			namsinh, err:= strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2023 - namsinh
			r := fmt.Sprintf("Tuoi cua ban la: %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}