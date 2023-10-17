package slashcommands

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/localization"
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
			Name:                     "admin",
			Description:              "admin_Description",
			NameLocalizations:        localization.LoadList("#admin"),
			DescriptionLocalizations: localization.LoadList("#admin.Description"),
			DMPermission:             &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:                     "allow-role",
					Description:              "allow-role_Description",
					NameLocalizations:        *localization.LoadList("#admin.allow-role"),
					DescriptionLocalizations: *localization.LoadList("#admin.allow-role.Description"),
					Type:                     discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:                     "value",
							Description:              "value_Description",
							NameLocalizations:        *localization.LoadList("#admin.allow-role.value"),
							DescriptionLocalizations: *localization.LoadList("#admin.allow-role.value.Description"),
							Type:                     discordgo.ApplicationCommandOptionBoolean,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								{
									Name:              "true",
									NameLocalizations: *localization.LoadList("#admin.allow-role.value.true"),
									Value:             true,
								},
								{
									Name:              "false",
									NameLocalizations: *localization.LoadList("#admin.allow-role.value.false"),
									Value:             false,
								},
							},
							Required: true,
						},
					},
				},
			},
			DefaultMemberPermissions: &DefaultMemberPermissions,
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			if len(options) == 0 {
				return
			}
			cmdName := options[0].Name

			switch cmdName {
			case "allow-role":
				admin_AllowRole(s, i, options[0].Options)
			}
		},
	})
}

func admin_AllowRole(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) {
	locale := i.Locale
	value := options[0].Value.(bool)

	guildDB, err := db.LoadGuildByID(i.GuildID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if guildDB.InProgress == true {
		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("`%s`", localization.Load(locale, "#admin.allow-role.InProgress")),
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		return
	}
	if value == true {
		p, err := s.UserChannelPermissions(s.State.User.ID, i.ChannelID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		if p&discordgo.PermissionManageRoles == 0 {
			err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
				Content:         fmt.Sprintf("`%s`", localization.Load(locale, "#admin.allow-role.NeedPermissions")),
				Flags:           discordgo.MessageFlagsEphemeral,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			}))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
			return
		}
	}
	if guildDB.AllowRole != value {
		db.UpdateGuildRoleOption(i.GuildID, value)

		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				builder.Embed().
					SetDescription(fmt.Sprintf("**%s** `%s` - %s",
						localization.Load(locale, "#admin.allow-role.proc.Title"),
						localization.Load(locale, "#admin.allow-role.proc.Description"),
						localization.Load(locale, "#admin.allow-role.proc.InProgress"),
					)).
					MessageEmbed,
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
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
					display, err := role.GetDisplay(member.User.ID)
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						continue
					}
					if value == true {
						err = role.Set(s, i.GuildID, member.User.ID, display)
					} else {
						err = role.Remove(s, i.GuildID, member.User.ID, display)
					}
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						continue
					}
				}
			}(res)

			after = res[len(res)-1].User.ID
		}
		wait.Wait()

		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				builder.Embed().
					SetDescription(fmt.Sprintf("**%s** `%s` - %s",
						localization.Load(locale, "#admin.allow-role.proc.Title"),
						localization.Load(locale, "#admin.allow-role.proc.Description"),
						localization.Load(locale, "#admin.allow-role.proc.Done"),
					)).
					MessageEmbed,
			},
		})
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		db.EndOFGuildProgress(i.GuildID)
	} else {
		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("`%s`", localization.Load(locale, "#admin.allow-role.Keep")),
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
	}
}
