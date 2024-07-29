package discord

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdown or not")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func Entry(token string) {
    // TODO: ensure other bots are not dependent on this one (and vice versa)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as: %v%v\n", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
        panic(fmt.Sprintf("Cannot open the session: %v", err))
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			panic(fmt.Sprintf("Cannot create '%v' command: %v", v.Name, err))
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
    fmt.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				panic(fmt.Sprintf("Cannot delete '%v' command: %v", v.Name, err))
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
