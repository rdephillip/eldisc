# eldisc - Donor Drive API Connected Discord Bot
This is a very simple Discord bot using DiscordGo which also communicates with your configured Donor Drive event participant page. I built this to work with Extra Life, but I've made it configurable. I have not tested this with any other events from Donor Drive at this time. Once I am able to do so, I will look at what is different between each connection to see if adjustments are necessary.

## Initial Configuration
When you first launch the executable it will attmept to detect a config.json file. Don't worry, it will create a default one and aske you for values if it doesn't exist. I will briefly explain the values that will be asked for and how you should fill them out.
Please note that I will explain configuring for Extra Life participants. If you have a different Donor Drive API you are intending to connect to know that I have not yet tested any others, and you may need to modify the Go to point to the appropriate endpoints.

- FirstSetup: This is a boolean value (true/false) which the program uses and stores to either initialize the configuration file or to load your custom settings
- Token: This is your Discord bot token. This is created when you configure your Discord Developer Portal application. Visit [Discord Developer Portal Documentation](https://discord.com/developers/docs/intro) for guidance from Discord on this process.
- Url: This is the API URL you intend to access. For Extra Life you would use https://www.extra-life.org/api but this will vary depending on the API you're accessing.
- Pid: This is the Participant ID. Extra Life will have this as a number in the URL of your page on their site.
- AdminAuthor: This is the Discord user ID which is assigned rights to use the !config commands and that can bypass the timer for the status command. This should be set to you. In Discord click your profile image in the lower right, then select Copy User ID.
- ChannelID: This is the channel you want the bot to announce its messages to. You can get this by right clicking the channel and clicking the Copy Channel ID. After the bots first setup is complete, you can also use **!config set channelid this** in the channel you want to set it to.
- Command: This is the command that allows users to check the goal status. The default is **!status**, but you can set it to whatever you choose. I used **!elstatus**
- Charity: This is the name of the charity you're connecting to. This is to allow the messages to be dynamically updated to the correct name if you plan to frequently adapt your bot to your needs.
