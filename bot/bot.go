package bot

import (
	"bot/environment"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// OriginalChannelID is saved at the start of the poll creation to return poll in the channel where it was initialized
type Poll struct {
	OriginalChannelID string
	Question          string
	Option1           string
	Option2           string
	Option3           string
	Option4           string
}

// Map of polls from different people, to handle multiple poll creation processes at once
// ChannelID of the Direct messages is used as unique key in the map
var responses map[string]Poll = map[string]Poll{}

// Returning embed format of the poll for visual representation in message
func (poll *Poll) toMessageEmbed() *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Option 1",
			Value: poll.Option1,
		},
		{
			Name:  "Option 2",
			Value: poll.Option2,
		},
	}
	if poll.Option3 != "" {
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Option 3", Value: poll.Option3})
	}
	if poll.Option4 != "" {
		fields = append(fields, &discordgo.MessageEmbedField{Name: "Option 4", Value: poll.Option4})
	}
	embed := discordgo.MessageEmbed{
		Title:  poll.Question,
		Fields: fields,
	}
	return &embed
}

func isNumeric(input string) bool {
	_, err := strconv.ParseInt(input, 10, 64)
	return err == nil
}

func Start() {
	sess, err := discordgo.New("Bot " + environment.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Main message handler
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignoring this bot's messages
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.GuildID == "" {
			directMessageHandler(s, m)
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != environment.BotPrefix || len(args) < 2 {
			return
		}

		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "World!")
			return
		}

		if args[1] == "poll" {
			PollCreation(s, m)
			return
		}

		if args[1] == "dividers" {
			// Creating a thread for hard calculation to not stop working
			go findDividersHandler(s, m, args)
			return
		}

		if args[1] == "help" {
			helpHandler(s, m)
			return
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Connection to discord bot session
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	log.Println("Bot is online!")

	// Working until termination signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
