package dcli

import (
	"fmt"
	"log"
	"strings"
	"time"

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

	fmt.Println(formatManicMinuteWord(state, time.Now().UTC()))
}

func formatManicMinuteWord(state *db.ManicMinuteRuntimeState, now time.Time) string {
	cooldownLabel := "none"
	if state != nil && state.HasCooldown && !state.CooldownUntil.IsZero() {
		remaining := state.CooldownUntil.Sub(now).Round(time.Second)
		if remaining > 0 {
			cooldownLabel = remaining.String()
		}
	}

	return fmt.Sprintf(
		"Current manic-word: %s (current chance: %.0f%%, cooldown remaining: %s)",
		state.TriggerWord,
		state.StartChance*100,
		cooldownLabel,
	)
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
