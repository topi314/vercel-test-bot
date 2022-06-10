package commands

import "github.com/disgoorg/disgo/discord"

func PingCommandHandler(interaction discord.ApplicationCommandInteraction) discord.InteractionResponse {
	return discord.InteractionResponse{
		Type: discord.InteractionCallbackTypeCreateMessage,
		Data: discord.MessageCreate{
			Content: "Pong!",
		},
	}
}
