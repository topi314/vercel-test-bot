package commands

import "github.com/disgoorg/disgo/discord"

func PingCommandHandler(interaction discord.ApplicationCommandInteraction, done func()) discord.InteractionResponse {
	defer done()
	return discord.InteractionResponse{
		Type: discord.InteractionResponseTypeCreateMessage,
		Data: discord.MessageCreate{
			Content: "Pong!",
		},
	}
}
