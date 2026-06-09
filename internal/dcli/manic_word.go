package dcli

import (
	"fmt"
	"log"
	"strings"

	"gitlab.int.magneato.site/dungar/prototype/internal/db"
	"gitlab.int.magneato.site/dungar/prototype/internal/utils"
)

// PrintManicMinuteWord prints the currently armed manic-minute trigger word for
// the configured protocol mode.
func PrintManicMinuteWord() {
	utils.LoadSettingsAndSecrets()
	db.ConnectToDatabase()

	mode, err := normalizeManicMinuteProtocolMode(utils.ProtocolMode())
	if err != nil {
		log.Fatal(err)
	}

	state, err := db.GetManicMinuteRuntimeState(mode)
	if err != nil {
		log.Fatal(err)
	}

	if state == nil || strings.TrimSpace(state.TriggerWord) == "" {
		log.Fatalf("no persisted manic minute trigger word for mode=%s; start dungar run first", mode)
	}

	fmt.Printf("Current manic-word: %s (current chance: %.0f%%)\n", state.TriggerWord, state.StartChance*100)
}

func normalizeManicMinuteProtocolMode(mode string) (string, error) {
	switch {
	case strings.EqualFold(mode, "discord"):
		return "discord", nil
	case strings.EqualFold(mode, "slack"):
		return "slack", nil
	default:
		return "", fmt.Errorf("unknown protocol mode for manic-word: %s", mode)
	}
}
