package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ReactionCounter(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	memberInfo := HandleEvent(s, r.Member, r.GuildID)
	memberInfo.Reactions++
	memberInfo.InfoChanged = true

	log.Println("Reactions for member", memberInfo.UserID, "is now", memberInfo.Reactions)
}
