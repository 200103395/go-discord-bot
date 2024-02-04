package bot

import (
	"bot/environment"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"time"
)

// Handling direct messages from users, in this case only poll creation process
func directMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	answers, ok := responses[m.ChannelID]
	if !ok {
		return
	}
	// If the user wants to create poll with less options
	if answers.Question != "" && answers.Option1 != "" && answers.Option2 != "" && m.Content == "Create Poll" {
		createAPoll(s, m)
		return
	}
	// Handling every situation
	if answers.Question == "" {
		answers.Question = m.Content
		responses[m.ChannelID] = answers

		s.ChannelMessageSend(m.ChannelID, "Great! What would be an Option 1?")
	} else if answers.Option1 == "" {
		answers.Option1 = m.Content
		responses[m.ChannelID] = answers

		s.ChannelMessageSend(m.ChannelID, "Great! What would be an Option 2?")
	} else if answers.Option2 == "" {
		answers.Option2 = m.Content
		responses[m.ChannelID] = answers

		s.ChannelMessageSend(m.ChannelID, "Great! What would be an Option 3?")
		s.ChannelMessageSend(m.ChannelID, "If you want to create the poll with 2 option type \"Create Poll\"")
	} else if answers.Option3 == "" {
		answers.Option3 = m.Content
		responses[m.ChannelID] = answers

		s.ChannelMessageSend(m.ChannelID, "Great! What would be an Option 4?")
		s.ChannelMessageSend(m.ChannelID, "If you want to create the poll with 3 option type \"Create Poll\"")
	} else {
		answers.Option4 = m.Content
		responses[m.ChannelID] = answers

		createAPoll(s, m)
	}
}

func helpHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	fields := []*discordgo.MessageEmbedField{
		{
			Name:  environment.BotPrefix + " help",
			Value: "Returning all features as a response",
		},
		{
			Name:  environment.BotPrefix + " hello",
			Value: "Returning World! as a response",
		},
		{
			Name:  environment.BotPrefix + " poll",
			Value: "Creating a poll",
		},
		{
			Name:  environment.BotPrefix + " dividers X",
			Value: "Finding all dividers for the number X",
		},
	}

	embed := discordgo.MessageEmbed{
		Title:  "Discord bot featuring ChatGPT",
		Fields: fields,
	}

	// Returning embed message of every possible command
	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

// Function that creates a poll from given information
func createAPoll(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := responses[m.ChannelID]; err == false {
		return
	}
	answers := responses[m.ChannelID]
	pollMessage := answers.toMessageEmbed()

	// Sending the poll message
	msg, err := s.ChannelMessageSendEmbed(answers.OriginalChannelID, pollMessage)
	if err != nil {
		fmt.Println("Error sending poll message:", err)
		return
	}

	// Adding reaction for the poll (numbers from 1 to 4)
	err = s.MessageReactionAdd(answers.OriginalChannelID, msg.ID, "1⃣")
	if err != nil {
		log.Println(err)
	}
	s.MessageReactionAdd(answers.OriginalChannelID, msg.ID, "2⃣")
	if answers.Option3 != "" {
		s.MessageReactionAdd(answers.OriginalChannelID, msg.ID, "3⃣")
	}
	if answers.Option4 != "" {
		s.MessageReactionAdd(answers.OriginalChannelID, msg.ID, "4⃣")
	}

	// At the end clearing map from this poll
	delete(responses, m.ChannelID)
}

// Initializing poll creation process
func PollCreation(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Panic(err)
	}

	// If the user initializing poll currently not creating any other poll, start a new process
	if _, ok := responses[channel.ID]; !ok {
		responses[channel.ID] = Poll{
			OriginalChannelID: m.ChannelID,
			Question:          "",
			Option1:           "",
			Option2:           "",
			Option3:           "",
			Option4:           "",
		}
		s.ChannelMessageSend(channel.ID, "Hey there, let's create a poll!")
		s.ChannelMessageSend(channel.ID, "What would be the question")
	} else {
		s.ChannelMessageSend(channel.ID, "Hey there, we are still creating a poll :smile")
	}
}

// Function to find every possible dividers for a given number
func findDividersHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Please, provide a number!")
		return
	}
	if isNumeric(args[2]) != true {
		s.ChannelMessageSend(m.ChannelID, "Please, use standard number consisting of digit")
		return
	}
	num, _ := strconv.ParseInt(args[2], 10, 64)
	var i int64
	result := "Dividers for " + args[2] + " are: "
	for i = 1; i*i <= num; i++ {
		if num%i == 0 {
			result += strconv.FormatInt(i, 10) + ", "
			if i != num/i {
				result += strconv.FormatInt(num/i, 10) + ", "
			}
		}
	}
	time.Sleep(10 * time.Second)
	s.ChannelMessageSend(m.ChannelID, result)
}
