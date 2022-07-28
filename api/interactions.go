package api

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/log"
	"vercel-test-bot/commands"
)

var handlers = map[string]func(client rest.Rest, interaction discord.ApplicationCommandInteraction) discord.InteractionResponse{
	"ping": commands.PingCommandHandler,
}

func HandleInteractions(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()
	logger.SetLevel(log.LevelTrace)
	client := rest.New(rest.NewClient(""))

	publicKey, err := hex.DecodeString(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpserver.HandleInteraction(publicKey, logger, func(respondFunc httpserver.RespondFunc, event httpserver.EventInteractionCreate) {
		switch i := event.Interaction.(type) {
		case discord.PingInteraction:
			err = respondFunc(discord.InteractionResponse{
				Type: discord.InteractionResponseTypePong,
			})

		case discord.ApplicationCommandInteraction:
			handler, ok := handlers[i.Data.CommandName()]
			if !ok {
				err = fmt.Errorf("command %s not implemented", i.Data.CommandName())
				break
			}
			err = respondFunc(handler(client, i))
		}
		if err != nil {
			logger.Errorf("error while handling interaction: %s", err)
		}
	})(w, r)
}
