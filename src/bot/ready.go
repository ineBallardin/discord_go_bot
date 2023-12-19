package ready

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ineBallardin/discord_go_bot/src/commands"
	"github.com/ineBallardin/discord_go_bot/src/events"
	"github.com/joho/godotenv"
)

var (
	Token   string
	guildID string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("TOKEN")
	guildID = os.Getenv("GUILD_ID")
}

func Bot() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent
	commands.RegisterSlashCommands(dg, guildID)
	events.RegisterEvents(dg)

	ticker := time.NewTicker(24 * time.Minute)

	go func() {
		for range ticker.C {
			events.SendMessageCounts(dg)
			events.SendReactionCounts(dg)
			events.SendVoiceChannelCounts(dg)
		}
	}()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
