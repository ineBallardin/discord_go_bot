package events

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MemberInfo struct {
	UserID              string
	IsImpulserPro       bool
	IsImpulser          bool
	Messages            int
	Reactions           int
	VoiceChannels       map[string]int64
	CurrentVoiceChannel string
	ChannelID           string
	JoinTime            time.Time
	LeaveTime           time.Time
	TotalTime           time.Duration
	InfoChanged         bool
}

var memberInfos = make([]*MemberInfo, 0)

func HandleEvent(s *discordgo.Session, member *discordgo.Member, guildID string) *MemberInfo {
	isImpulserPro := false
	isImpulser := false
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			log.Println("Erro ao obter cargo:", err)
			continue
		}
		if role.Name == "impulserPRO" {
			isImpulserPro = true
			break
		}
		if role.Name == "impulser" {
			isImpulser = true
			break
		}
	}

	memberInfo, exists := findMemberInfo(member.User.ID)
	if !exists {
		memberInfo = &MemberInfo{
			UserID:        member.User.ID,
			IsImpulserPro: isImpulserPro,
			IsImpulser:    isImpulser,
			Messages:      0,
			Reactions:     0,
			VoiceChannels: make(map[string]int64),
		}
		memberInfos = append(memberInfos, memberInfo)
	}
	return memberInfo
}

func findMemberInfo(userID string) (*MemberInfo, bool) {
	for _, memberInfo := range memberInfos {
		if memberInfo.UserID == userID {
			return memberInfo, true
		}
	}
	return nil, false
}

func FormatMemberInfo(s *discordgo.Session, userID string, guildID string) string {
	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		log.Println("Erro ao obter membro do servidor:", err)
		return ""
	}

	memberInfo := HandleEvent(s, member, guildID)

	if memberInfo.InfoChanged {
		log.Println("As informações dos membros foram atualizadas")
		memberInfo.InfoChanged = false
		return fmt.Sprintf(
			"## Métricas do dia %s\n- Membro: <@%s>\n    - IsImpulserPro: %v\n    - IsImpulser: %v\n    - Mensagens enviadas: %d\n    - Reações: %d\n    - Tempo em canal de voz: %.2f",
			time.Now().Format("02/01/2006"),
			memberInfo.UserID,
			memberInfo.IsImpulserPro,
			memberInfo.IsImpulser,
			memberInfo.Messages,
			memberInfo.Reactions,
			float64(memberInfo.VoiceChannels["totalTime"])/float64(time.Minute),
		)
	}

	return ""
}

func GetMemberInfos() []*MemberInfo {
	return memberInfos
}
