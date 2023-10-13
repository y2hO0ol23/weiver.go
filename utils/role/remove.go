package role

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/utils/database"
)

func Remove(s *discordgo.Session, guildID string, memberID string, display string) {
	roleDB := db.GetRoleByInfo(guildID, display)
	if roleDB == nil {
		return
	}

	err := s.GuildMemberRoleRemove(guildID, memberID, roleDB.RoleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok := db.RemoveRoleOnUser(roleDB.ID, memberID); !ok {
		err := s.GuildRoleDelete(guildID, roleDB.RoleID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}