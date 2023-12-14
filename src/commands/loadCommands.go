package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ineBallardin/discord_go_bot/src/bot"
)

func LoadCommands(bot *bot.Bot) {
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
