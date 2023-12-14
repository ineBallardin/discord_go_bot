package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("TOKEN")
}

func main() {
	bot := createBot()
	dg := createDiscordSession(bot)
	loadCommands(bot)
	loadEvents("./src/events", bot)
	openConnection(dg)
	addHandler(dg)
	waitForExit()
}

func createBot() *Bot {
	return &Bot{
		Cooldowns: make(map[string]string),
		Commands:  make(map[string]Command),
	}
}

func createDiscordSession(bot *Bot) *discordgo.Session {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		os.Exit(1)
	}
	bot.Session = dg
	return dg
}

func loadCommands(bot *Bot) {
	root := "./src/commands"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			content, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v", path, err)
				return nil
			}

			// carregar os comandos aqui
			fmt.Println(path)
			fmt.Println(string(content))
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}
}

func openConnection(dg *discordgo.Session) {
	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		os.Exit(1)
	}
}

func addHandler(dg *discordgo.Session) {
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot", r.User.Username, "is now running.")
	})
}

func waitForExit() {
	<-make(chan struct{})
}

func loadEvents(root string, bot *Bot) {
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			// Carregar os eventos aqui
			fmt.Println(file.Name())
		}
	}
}
