package events

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type VoiceChannelStats struct {
	UserID        string
	Roles         []string
	ChannelID     string
	JoinTime      time.Time
	LeaveTime     time.Time
	TotalTime     time.Duration
	IsImpulserPro bool
}

var voiceChannelStats = make(map[string]*VoiceChannelStats)

func VoiceChannelCounter(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	stats, ok := voiceChannelStats[v.UserID]
	if !ok {
		stats = &VoiceChannelStats{
			UserID:    v.UserID,
			ChannelID: v.ChannelID,
			JoinTime:  time.Now(),
		}
		voiceChannelStats[v.UserID] = stats
	} else {
		stats.LeaveTime = time.Now()
		stats.TotalTime = stats.LeaveTime.Sub(stats.JoinTime)
		if stats.TotalTime.Minutes() >= 5 {
			member, err := s.GuildMember(v.GuildID, v.UserID)
			if err != nil {
				log.Println("Erro ao obter membro do servidor:", err)
				return
			}
			stats.Roles = make([]string, len(member.Roles))
			for i, roleID := range member.Roles {
				role, err := s.State.Role(v.GuildID, roleID)
				if err != nil {
					log.Println("Erro ao obter cargo:", err)
				}
				stats.Roles[i] = role.Name
			}

			isImpulserPro := false
			for _, role := range stats.Roles {
				if role == "impulserPRO" {
					isImpulserPro = true
					break
				}
			}
			stats.IsImpulserPro = isImpulserPro
		}
		stats.JoinTime = time.Now()
		stats.ChannelID = v.ChannelID
	}
}

func SendVoiceChannelCounts(s *discordgo.Session) {
	for _, stats := range voiceChannelStats {
		totalTimeInMinutes := stats.TotalTime.Minutes()
		formattedTotalTime := fmt.Sprintf("%.2f", totalTimeInMinutes)
		if totalTimeInMinutes >= 5 {
			_, err := s.ChannelMessageSend("1101510837555974176", fmt.Sprintf("## Contador de Canais de Voz\n**Membro:** <@%s>,\n- **impulserPRO:** %v\n- Canal: <#%s>\n- Tempo total: %s minutos", stats.UserID, stats.IsImpulserPro, stats.ChannelID, formattedTotalTime))
			if err != nil {
				log.Println("Erro ao enviar mensagem para o canal:", err)
			}
		}
	}
}
