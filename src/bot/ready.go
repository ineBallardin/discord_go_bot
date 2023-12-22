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
	"github.com/ineBallardin/discord_go_bot/src/infra"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
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

	infra.InitializeAppWithServiceAccount()

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent
	commands.RegisterSlashCommands(dg, guildID)
	events.RegisterEvents(dg)

	// ticker := time.NewTicker(1 * time.Minute)

	// go func() {
	// 	for range ticker.C {
	// 		SendUpdates(dg, guildID)
	// 	}
	// }()

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 * * *", func() { SendUpdates(dg, guildID) })
	c.Start()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func SendUpdates(s *discordgo.Session, guildID string) {
	var messages string
	for _, memberInfo := range events.GetMemberInfos() {
		if memberInfo.InfoChanged {
			minutes := memberInfo.TotalTime / time.Minute
			seconds := (memberInfo.TotalTime % time.Minute) / time.Second
			totalTimeInVoiceChannel := fmt.Sprintf("%d.%02ds", minutes, seconds)
			message := fmt.Sprintf(
				"## Membro: <@%s>\n    - IsImpulserPro: %v\n    - IsImpulser: %v\n    - Mensagens enviadas: %d\n    - Reações: %d\n    - Tempo em canal de voz: %s\n",
				memberInfo.UserID,
				memberInfo.IsImpulserPro,
				memberInfo.IsImpulser,
				memberInfo.Messages,
				memberInfo.Reactions,
				totalTimeInVoiceChannel,
			)
			messages += message
			memberInfo.InfoChanged = false
		}
	}
	if messages != "" {
		title := fmt.Sprintf("# Métricas do dia %s\n", time.Now().Format("02/01/2006"))
		_, err := s.ChannelMessageSend("1101510837555974176", title+messages)
		if err != nil {
			log.Println("Erro ao enviar mensagem para o canal:", err)
		}
		ResetMemberInfos()
	}
}

func ResetMemberInfos() {
	for _, memberInfo := range events.GetMemberInfos() {
		memberInfo.IsImpulserPro = false
		memberInfo.IsImpulser = false
		memberInfo.Messages = 0
		memberInfo.Reactions = 0
		memberInfo.VoiceChannels = make(map[string]int64)
		memberInfo.CurrentVoiceChannel = ""
		memberInfo.ChannelID = ""
		memberInfo.JoinTime = time.Time{}
		memberInfo.LeaveTime = time.Time{}
		memberInfo.TotalTime = 0
		memberInfo.InfoChanged = false
	}
}
