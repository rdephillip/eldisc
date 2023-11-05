package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

func loadConfig() {
	config, err := os.ReadFile("./config.json") // Try to read config.json
	// If error, check if file exists, if not initiate first setup sequence
	if err != nil {
		if os.IsNotExist(err) {
			FirstSetup = true
			saveConfig(FirstSetup)
		} else {
			log.Println(err.Error())
		}
	} else {
		// If no error pull first setup to verify file is configured
		FirstSetup, err = strconv.ParseBool(gjson.Get(string(config), "FirstSetup").String())
		if err != nil {
			log.Println(err.Error())
		}
	}

	// If this is the first setup, request user input to setup bot
	if FirstSetup {
		log.Println("First run detected. Please enter requested information.")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Bot Token (Discord bot Auth Token) >> ")
		Token, _ = reader.ReadString('\n')
		Token = strings.TrimSpace(Token)

		fmt.Print("Enter API Url (Donor Drive) >> ")
		Url, _ = reader.ReadString('\n')
		Url = strings.TrimSpace(Url)

		fmt.Print("Enter Extra Life Participant ID (Numerical ID number) >> ")
		Pid, _ = reader.ReadString('\n')
		Pid = strings.TrimSpace(Pid)

		fmt.Print("Enter Admin Author (Discord user ID number) >> ")
		AdminAuthor, _ = reader.ReadString('\n')
		AdminAuthor = strings.TrimSpace(AdminAuthor)

		fmt.Print("Enter Channel ID (Discord channel ID number for announcing) >> ")
		ChannelID, _ = reader.ReadString('\n')
		ChannelID = strings.TrimSpace(ChannelID)

		fmt.Print("Enter Charity Name >> ")
		Charity, _ = reader.ReadString('\n')
		Charity = strings.TrimSpace(Charity)

		FirstSetup = false
	} else {
		// If it isn't the first setup, load the configuration file
		Token = gjson.Get(string(config), "Token").String()
		Url = gjson.Get(string(config), "Url").String()
		Pid = gjson.Get(string(config), "Pid").String()
		AdminAuthor = gjson.Get(string(config), "AdminAuthor").String()
		ChannelID = gjson.Get(string(config), "ChannelID").String()
		Command = gjson.Get(string(config), "Command").String()
		Charity = gjson.Get(string(config), "Charity").String()
	}

	log.Println("Configuration file loaded.")
}

func saveConfig(makeFile bool) {
	// Initialize empty string for JSON data
	saveData := ""

	// If this is the first setup, create the default empty data
	if makeFile {
		saveData = "{\n" +
			"\t\"FirstSetup\":\"true\",\n" +
			"\t\"Token\":\"\",\n" +
			"\t\"Url\":\"\",\n" +
			"\t\"Pid\":\"\",\n" +
			"\t\"AdminAuthor\":\"\",\n" +
			"\t\"ChannelID\":\"\",\n" +
			"\t\"Command\":\"!status\"\n" +
			"\t\"Charity\":\"\"\n" +
			"}"
	} else {
		// If not first setup, grab global variables and prepare JSON string
		saveData = fmt.Sprintf("{\n"+
			"\t\"FirstSetup\":\"%s\",\n"+
			"\t\"Token\":\"%s\",\n"+
			"\t\"Url\":\"%s\",\n"+
			"\t\"Pid\":\"%s\",\n"+
			"\t\"AdminAuthor\":\"%s\",\n"+
			"\t\"ChannelID\":\"%s\",\n"+
			"\t\"Command\":\"%s\"\n"+
			"\t\"Charity\":\"%s\"\n"+
			"}", strconv.FormatBool(FirstSetup), Token, Url, Pid, AdminAuthor, ChannelID, Command, Charity)
	}

	// Open file in truncate mode to overwrite existing configuration with current configuration variables and write to the file
	file, err := os.OpenFile("./config.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println(err.Error())
	}

	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Println(err.Error())
	}

	file.Write([]byte(saveData))
	log.Println("Configuration file saved.")
}
