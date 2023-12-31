package events

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func VoiceChannelCounter(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	member, err := s.GuildMember(v.GuildID, v.UserID)
	if err != nil {
		log.Println("Erro ao obter membro do servidor:", err)
		return
	}

	memberInfo := HandleEvent(s, member, v.GuildID)

	if memberInfo.ChannelID == "" {
		memberInfo.ChannelID = v.ChannelID
		memberInfo.JoinTime = time.Now()
		log.Println("Membro entrou às", memberInfo.JoinTime)
	} else {
		timeInChannel := time.Since(memberInfo.JoinTime)
		if timeInChannel.Minutes() >= 5 {
			memberInfo.LeaveTime = time.Now()
			memberInfo.TotalTime += timeInChannel
			memberInfo.InfoChanged = true
		}

	}
}
