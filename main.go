package main

import (
	"fmt"
	"hush-bot/data"
	"hush-bot/discord"
	"hush-bot/slack"
	"hush-bot/teams"
	"hush-bot/utils"
)

func main() {
	dt := utils.GetDockerSecret("discord_token")
	st := utils.GetDockerSecret("slack_token")
	tt := utils.GetDockerSecret("teams_token")

    fmt.Println("Initializing database...")
	ftl := data.InitDbForUserClient()
	if ftl != nil {
		fmt.Println(ftl)
		panic("Fatal initialization error. Shutting down. ")
	}

	if dt != "" {
		fmt.Println("Spawning Discord Bot...")
		go discord.Entry(dt)
	}

	if st != "" {
		fmt.Println("Spawning Slack Bot...")
		go slack.Entry(st)
	}

	if tt != "" {
		fmt.Println("Spawning Teams Bot...")
		go teams.Entry(tt)
	}
}
