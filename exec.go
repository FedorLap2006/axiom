package main

import (
	api "botapi"
	dsgo "github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func test(event *api.Event) {
	msg := event.Data.(*dsgo.Message)
	log.Println(msg.Content)
}

func test2(event *api.Event) {
	chc := event.Data.(*dsgo.Channel)
	log.Println(chc.Name)
}

func main() {
	bot := api.Bot{}
	if len(os.Args) > 1 {
		bot.Init(os.Args[1])
	} else {
		bot.Init(os.Getenv("DISCORD_BOTAPI_TOKEN"))
	}
	api.Utils_PREFIX = "^!"
	api.Utils_CMDS = map[string]api.EventHandler{
		"hi my friend": test,
	}
	bot.Handlers = api.HandlersList{
		api.EVENTH_MESSAGE_CREATE:    api.Utils_CMDSHandler,
		api.EVENTH_GUILD_CHAN_CREATE: test2,
	}
	e := bot.Run()
	if e != nil {
		log.Fatal(e)
	}
	log.Printf(`Discord %s running`, bot.Botu.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	e = bot.Stop()
	if e != nil {
		log.Fatal(e)
	}
}
