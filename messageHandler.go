package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// If the message author is the bot, ignore the message
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the configured command is used and it has been either more than 5 minutes or this is the Admin
	if m.Content == Command && (time.Since(LastStatusCheck) > 300000000000 || m.Author.ID == AdminAuthor) {
		// Retrieve participant info
		goal, raised, avatar, donate := participantInfo()

		// Discord embed for more attractiv, and noticeable response
		embed := &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s Goal Status", Charity),
			URL:         donate,
			Description: fmt.Sprintf("$%s of $%s raised.", raised, goal),
			Color:       0xDB842B,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: avatar,
			},
		}

		// Send embed message and report to console who used the command
		s.ChannelMessageSendEmbed(ChannelID, embed)
		log.Println(Command + " command used by: " + m.Author.Username)

		// If it is the Admin, don't update the timer
		if m.Author.ID != AdminAuthor {
			LastStatusCheck = time.Now()
		}
	}

	// If the messages are from the Admin process for acceptable configuration commands
	if m.Author.ID == AdminAuthor {
		com := strings.Split(m.Content, " ")

		if len(com) > 1 {
			if com[0] == "!config" {
				switch len(com) {
				case 4:
					switch strings.ToLower(com[1]) {
					case "set":
						switch strings.ToLower(com[2]) {
						case "token":
							Token = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "Bot token set.")
						case "url":
							Url = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "API Url set.")
						case "pid":
							Pid = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "Participant id set.")
						case "adminauthor":
							AdminAuthor = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "Admin author set. (If you lose control edit the config.json with the correct id and restart the bot)")
						case "channelid":
							if strings.ToLower(com[3]) == "this" {
								ChannelID = m.ChannelID
								saveConfig(false)
							} else {
								ChannelID = com[3]
								saveConfig(false)
							}
							s.ChannelMessageSend(m.ChannelID, "Announcement channel set.")
						case "command":
							Command = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "Status command set.")
						case "charity":
							Charity = com[3]
							saveConfig(false)
							s.ChannelMessageSend(m.ChannelID, "Charity name set.")
						}
					}
				default:
					s.ChannelMessageSend(m.ChannelID, "Invalid admin command")
				}
			}
		} else if com[0] == "!test" {
			testDono(s)
		}
	}
}

func testDono(s *discordgo.Session) {
	var raisedVal = 50.00
	var amountVal = 50.00

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s Donation", "Test Charity"),
		URL:         "",
		Description: fmt.Sprintf("$%.2f donation received.", amountVal),
		Color:       0xDB842B,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Test Donation",
				Value:  "This is not a real donation. This is just a test.",
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("%s Goal Status", "Test Charity"),
				Value:  fmt.Sprintf("$%.2f of $%s", raisedVal+amountVal, "1000"),
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendEmbed(ChannelID, embed)
}
