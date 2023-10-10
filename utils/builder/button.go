package builder

import "github.com/bwmarrin/discordgo"

type ButtonStructure struct {
	discordgo.Button
}

func Button() *ButtonStructure {
	return &ButtonStructure{}
}

func (b *ButtonStructure) SetCustomId(value string) *ButtonStructure {
	b.CustomID = value
	return b
}

func (b *ButtonStructure) SetLable(value string) *ButtonStructure {
	b.Label = value
	return b
}

func (b *ButtonStructure) SetStyle(value discordgo.ButtonStyle) *ButtonStructure {
	b.Style = value
	return b
}
