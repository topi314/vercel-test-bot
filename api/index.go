package api

import (
	"encoding/hex"
	"net/http"
	"os"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/httpserver"
	"github.com/disgoorg/disgo/json"
	"vercel-test-bot/commands"
)

var PublicKey = os.Getenv("PUBLIC_KEY")

var handlers = []func(interaction discord.ApplicationCommandInteraction) discord.InteractionResponse{
	commands.PingCommandHandler,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	publicKey, err := hex.DecodeString(PublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("Public key: " + string(publicKey))

	if !httpserver.VerifyRequest(r, publicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	var interaction discord.UnmarshalInteraction
	if err = json.NewDecoder(r.Body).Decode(&interaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var response discord.InteractionResponse
	switch i := interaction.Interaction.(type) {
	case discord.PingInteraction:
		response = discord.InteractionResponse{
			Type: discord.InteractionCallbackTypePong,
		}

	case discord.ApplicationCommandInteraction:
		for _, handler := range handlers {
			response = handler(i)
			break
		}
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
