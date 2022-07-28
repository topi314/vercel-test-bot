package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
)

func PingCommandHandler(client rest.Rest, interaction discord.ApplicationCommandInteraction) discord.InteractionResponse {
	return discord.InteractionResponse{
		Type: discord.InteractionResponseTypeCreateMessage,
		Data: discord.MessageCreate{
			Content: "Ping!",
		},
	}
}
