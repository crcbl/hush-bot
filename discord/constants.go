package discord

import (
	"github.com/bwmarrin/discordgo"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = true
	defaultMemberPermissions int64 = discordgo.PermissionManageMessages
	commands                       = []*discordgo.ApplicationCommand{
		{
			Name:                     "test",
			Description:              "Test command for demonstration of default command permissions",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
		},
		{
			Name:        "secret",
			Description: "Display a secret by name",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "secretName",
					Description: "Name of the secret which you are attempting to view",
					Required:    true,
				},
			},
		},
		{
			Name:        "permissions",
			Description: "Command to display names of secrets which have been shared with you, and number of times you may request to view said secret.",
		},
		{
			Name:        "share",
			Description: "Command to share a secret with a specified user, for a given number of uses and lifetime",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "secretName",
					Description: "Name of the secret which you are attempting to share",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "Tagged discord user of whom you would like to share this secret",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "-uses",
					Description: "Specify this subcommand if you would like to set a number of uses this secret is allowed. Lifetime is infinite if not specified.",
					Required:    false,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "count",
							Description: "Number of times you would like this user to be permitted to view this secret",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "-lifetime",
					Description: "Specify this command if you would like to set a lifetime on the secret sharing",
					Required:    false,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "hours",
							Description: "Number of hours you would like the secret permissions to expire after",
							MinValue:    &integerOptionMinValue,
						},
					},
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"test":        test,
		"secret":      secret,
		"permissions": permissions,
		"share":       share,
	}
)
