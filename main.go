// Code comes from https://github.com/bwmarrin/discordgo/blob/master/examples/pingpong/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Failed to create Discord session: ", err)
		return
	}

	// Register the onMessageCreate function as the callback for MessageCreate events
	dg.AddHandler(onMessageCreate)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Failed to open Discord connection: ", err)
		return
	}

	// Cleanly close down the Discord session after the main function finishes execution
	defer dg.Close()

	// Wait until CTRL-C or other termination signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigChan
}

// onMessageCreate will be called via the AddHandler function above every time
// a new message is created on any channel that the authenticated bot has
// access to.
func onMessageCreate(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	// This is not specifically required for this example, but is good practice.
	if msg.Author.ID == sess.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if msg.Content == "ping" {
		sess.ChannelMessageSend(msg.ChannelID, "Pong!")
	}
}
