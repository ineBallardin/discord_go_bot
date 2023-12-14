package events

import (
	"fmt"
	"os"
	"strings"

	"github.com/ineBallardin/discord_go_bot/src/bot"
)

func LoadEvents(root string, bot *bot.Bot) {
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
