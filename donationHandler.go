package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

// Function called every 15 seconds tracking a delta to see if any new donations arrived since last check.
// If they are found send embed with appropriate details
func checkDonations(s *discordgo.Session, t time.Time) {
	body := participantDonations()

	gjson.Parse(string(body)).ForEach(func(key, value gjson.Result) bool {
		layout := "2006-01-02T15:04:05-0700"
		donation, err := time.Parse(layout, gjson.Get(value.String(), "createdDateUTC").String())
		if err != nil {
			log.Println(err.Error())
		}

		if donation.After(LastCheck) {
			amount := gjson.Get(value.String(), "amount").String()
			displayName := gjson.Get(value.String(), "displayName").String()
			message := gjson.Get(value.String(), "message").String()

			// Request participant info
			goal, raised, avatar, donate := participantInfo()

			if displayName == "" {
				displayName = "Anonymous"
			}

			var raisedVal, _ = strconv.ParseFloat(raised, 64)

			embed := &discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s Donation", Charity),
				URL:         donate,
				Description: fmt.Sprintf("$%s donation received.", amount),
				Color:       0xDB842B,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: avatar,
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   displayName,
						Value:  message,
						Inline: true,
					},
					{
						Name:   fmt.Sprintf("%s Goal Status", Charity),
						Value:  fmt.Sprintf("$%.2f of $%s", raisedVal, goal),
						Inline: true,
					},
				},
			}

			s.ChannelMessageSendEmbed(ChannelID, embed)
		}

		return true
	})
}
