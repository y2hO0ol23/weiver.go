package slash_commands

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
	"github.com/y2hO0ol23/weiver/utils/role"
)

func init() {
	var (
		DMPermission             bool  = false
		DefaultMemberPermissions int64 = discordgo.PermissionManageRoles
	)

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "allow_role",
			Description:  "FOR ADMIN - DEFAULT:False",
			DMPermission: &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "value",
					Description: "Of role update",
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "true",
							Value: true,
						},
						{
							Name:  "false",
							Value: false,
						},
					},
					Required: true,
				},
			},
			DefaultMemberPermissions: &DefaultMemberPermissions,
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			value := options[0].Value.(bool)

			guildDB := db.LoadGuildByID(i.GuildID)
			if guildDB.AllowRole != value {
				db.UpdateGuildRoleOption(i.GuildID, value)

				err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						builder.Embed().
							SetDescription("**Update Option** `AllowRole` - in progress").
							MessageEmbed,
					},
				}))
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}

				var (
					after string
					wait  sync.WaitGroup
				)
				for {
					res, err := s.GuildMembers(i.GuildID, after, 1000)
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						continue
					}
					if len(res) == 0 {
						break
					}

					wait.Add(1)
					go func(members []*discordgo.Member) {
						defer wait.Done()
						for _, member := range members {
							if value == true {
								role.Set(s, i.GuildID, member.User.ID, role.GetDisplay(member.User.ID))
							} else {
								role.Remove(s, i.GuildID, member.User.ID, role.GetDisplay(member.User.ID))
							}
						}
					}(res)

					after = res[len(res)-1].User.ID
				}
				wait.Wait()

				_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{
						builder.Embed().
							SetDescription("**Update Option** `AllowRole` - in progress").
							MessageEmbed,
					},
				})
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
			} else {
				err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content: fmt.Sprintf("`Noting changes. AllowRole: %v", value),
					Flags:   discordgo.MessageFlagsEphemeral,
				}))
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
			}
		},
	})
}
