package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Quote struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

var (
	Token string
)

var quotes []Quote

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Read quotes on localFile
	// 	Unmarshal is the way that turn a JSON document into a Go struct.
	// Marshal is the opposite: we turn on Go struct to JSON document.
	jsonData, _ := ioutil.ReadFile("data.json")
	err := json.Unmarshal([]byte(jsonData), &quotes)

	rand.Seed(time.Now().Unix())

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session - ", err)
		return
	}

	for _, elem := range quotes {
		fmt.Println("Message: ", elem.Message)
		fmt.Println("Author: ", elem.Author)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is running. Use CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// Once any signal is received on sc, end the main function
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!quote" {
		n := rand.Int() % len(quotes)

		embed := &discordgo.MessageEmbed{
			Color:       0x8261DA,
			Description: quotes[n].Message,
			Title:       quotes[n].Author,
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			fmt.Println(err)
		}
	}

	if strings.HasPrefix(m.Content, "!addQuote") {
		messageString := m.Content[10:]
		fmt.Println(messageString)
		quotes = append(quotes, Quote{Message: messageString, Author: m.Author.Username})
		result, err := json.Marshal(quotes)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("data.json", result, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
