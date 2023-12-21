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
		memberInfo.LeaveTime = time.Now()
		memberInfo.TotalTime = memberInfo.LeaveTime.Sub(memberInfo.JoinTime)

		if memberInfo.TotalTime.Minutes() < 5 {
			return
		}

		log.Printf("Saiu às: %v\n Ficou %v", memberInfo.LeaveTime, memberInfo.TotalTime)

		memberInfo.JoinTime = time.Now()
		memberInfo.ChannelID = v.ChannelID
	}
}

// func SendVoiceChannelCounts(s *discordgo.Session) {
// 	for _, stats := range voiceChannelStats {
// 		totalTimeInMinutes := stats.TotalTime.Minutes()
// 		formattedTotalTime := fmt.Sprintf("%.2f", totalTimeInMinutes)
// 		if totalTimeInMinutes >= 5 {
// 			_, err := s.ChannelMessageSend("1101510837555974176", fmt.Sprintf("## Contador de Canais de Voz\n- **Membro:** <@%s>,\n- **impulserPRO:** %v\n- **impulser:** %v\n- Canal: <#%s>\n- Tempo total: %s minutos", stats.UserID, stats.IsImpulserPro, stats.IsImpulser, stats.ChannelID, formattedTotalTime))
// 			if err != nil {
// 				log.Println("Erro ao enviar mensagem para o canal:", err)
// 			}
// 		}
// 	}
// }
