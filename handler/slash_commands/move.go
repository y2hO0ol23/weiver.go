package slash_commands

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
	"github.com/y2hO0ol23/weiver/utils/builder"
	db "github.com/y2hO0ol23/weiver/utils/database"
	reviewutil "github.com/y2hO0ol23/weiver/utils/review"
)

func init() {
	var (
		DMPermission bool = false
	)

	commands = append(commands, form{
		data: &discordgo.ApplicationCommand{
			Name:         "move-review",
			Description:  "move review on this channel",
			DMPermission: &DMPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "subject",
					Description: "Select subject",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
		execute: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			fromID := i.Interaction.Member.User.ID
			toID := options[0].Value.(string)

			review, err := db.LoadReivewByInfo(fromID, toID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
			if review == nil {
				err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content:         "`No review on this subject`",
					Flags:           discordgo.MessageFlagsEphemeral,
					AllowedMentions: &discordgo.MessageAllowedMentions{},
				}))
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				return
			}

			subject, err := s.GuildMember(i.GuildID, toID)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			_, err = s.ChannelMessage(review.ChannelID, review.MessageID)
			if err != nil {
				review, err := reviewutil.Resend(s, i, review)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
				if review != nil {
					reviewutil.ModifyDM(s, review)
				}
				return
			}

			err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					reviewutil.EmbedBody(review, subject.AvatarURL("")).
						SetDescription(fmt.Sprintf("https://discord.com/channels/%s/%s/%s", review.GuildID, review.ChannelID, review.MessageID)).
						MessageEmbed,
				},
				Components: []discordgo.MessageComponent{
					builder.ActionRow().AddComponents(
						builder.Button().
							SetCustomID("move-review").
							SetLable("Move on this channel").
							SetStyle(discordgo.SuccessButton),
					).ActionsRow,
				},
				Flags:           discordgo.MessageFlagsEphemeral,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			}))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			msg, err := s.InteractionResponse(i.Interaction)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}

			var handler func(*discordgo.Session, *discordgo.InteractionCreate)
			handler = func(s *discordgo.Session, iter *discordgo.InteractionCreate) {
				if iter.Type != discordgo.InteractionMessageComponent || i.Interaction.Member.User.ID != iter.Interaction.Member.User.ID {
					s.AddHandlerOnce(handler)
					return
				}
				if iter.Interaction.Message.ID != msg.ID {
					return
				}

				data := iter.MessageComponentData()
				if data.ComponentType != discordgo.ButtonComponent || data.CustomID != "move-review" {
					return
				}

				s.InteractionResponseDelete(i.Interaction)
				reviewNow, err := db.LoadReivewByInfo(fromID, toID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					return
				}
				if review.TimeStamp != reviewNow.TimeStamp {
					err := s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							builder.Embed().
								SetDescription("❌ This review has been edited").
								MessageEmbed,
						},
					}))
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
					}
					return
				}
				err = reviewutil.DeleteMessage(s, fromID, toID)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				review, err = reviewutil.Resend(s, iter, review)
				if err != nil {
					log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				}
				if review != nil {
					reviewutil.ModifyDM(s, review)
				}
			}
			s.AddHandlerOnce(handler)
		},
	})
}
