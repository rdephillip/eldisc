// ===========================================================================================================================================
// Donor Drive API connected Discord bot
//		Version:		1.0.0
//		Author:			TheBarbarian_EL (R DePhillip)
//		Date:			2023-11-04
//		Description:	This is a simple bot that communicates with the Donor Drive API. It has only been tested against the EXTRA LIFE
//						JSON responses. Configuration is from a JSON file and will be initialized on first launch. Configuration commands
//						are available within the Discord environment.
// ===========================================================================================================================================
// TODO:
//		- ADD help command for config manual responses in Discord. Manual should be DM'd to Admin user on request.
// ===========================================================================================================================================

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Globals used throughout application, loaded from config.json
var (
	FirstSetup      bool
	Token           string
	Url             string
	Pid             string
	LastCheck       = time.Now()
	LastStatusCheck = time.Now()
	AdminAuthor     string
	ChannelID       string
	Command         string
	Charity         string
)

func main() {
	loadConfig()      // Run configuration load, if first run or file missing create file and request information
	saveConfig(false) // Force save configuration file

	// Create new discord session
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error creating Discord session,", err)
		return
	}

	// Identify message handler function
	dg.AddHandler(messageCreate)

	// Identify intents, we will only be using Guild Messages
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open connection
	err = dg.Open()
	if err != nil {
		log.Println("Error opening connection,", err)
	}

	// Display console information of bot status and hardcorded Discord configuration commands
	log.Println("Bot is now running. Press CTRL-C to exit.")
	log.Println("Default goal status command: !status")
	log.Println("Discord config commands:")
	log.Println("\t!config set token [token]")
	log.Println("\t!config set url [url]")
	log.Println("\t!config set pid [particpant id]")
	log.Println("\t!config set adminauthor [author ID]")
	log.Println("\t!config set channelid [channel ID/this]")
	log.Println("\t!config set command [new command]")
	log.Println("\t!config set charity [charity name]")

	// Identify 15 second timer to limit API calls and check for donations
	ticker := time.NewTicker(15 * 1000 * time.Millisecond)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	for {
		select {
		case <-sc:
			return
		case t := <-ticker.C:
			checkDonations(dg, t)
			LastCheck = time.Now()
		}
	}

	// Cleanup and save configuration on close
	ticker.Stop()
	dg.Close()
	saveConfig(false)
}
