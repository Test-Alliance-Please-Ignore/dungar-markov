package dcli

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.int.magneato.site/dungar/prototype/internal/db"
)

func TestNormalizeManicMinuteProtocolMode(t *testing.T) {
	mode, err := normalizeManicMinuteProtocolMode("discord")
	assert.NoError(t, err)
	assert.Equal(t, "discord", mode)

	mode, err = normalizeManicMinuteProtocolMode("Slack")
	assert.NoError(t, err)
	assert.Equal(t, "slack", mode)
}

func TestNormalizeManicMinuteProtocolModeUnknown(t *testing.T) {
	mode, err := normalizeManicMinuteProtocolMode("xmpp")
	assert.Error(t, err)
	assert.Equal(t, "", mode)
}

func TestFormatManicMinuteWordWithCooldown(t *testing.T) {
	now := time.Date(2026, time.June, 9, 16, 30, 0, 0, time.UTC)
	state := &db.ManicMinuteRuntimeState{
		TriggerWord:   "week",
		StartChance:   0.26,
		HasCooldown:   true,
		CooldownUntil: now.Add(17*time.Minute + 30*time.Second),
	}

	got := formatManicMinuteWord(state, now)
	assert.Equal(t, "Current manic-word: week (current chance: 26%, cooldown remaining: 17m30s)", got)
}

func TestFormatManicMinuteWordWithoutCooldown(t *testing.T) {
	now := time.Date(2026, time.June, 9, 16, 30, 0, 0, time.UTC)
	state := &db.ManicMinuteRuntimeState{
		TriggerWord: "week",
		StartChance: 0.20,
	}

	got := formatManicMinuteWord(state, now)
	assert.Equal(t, "Current manic-word: week (current chance: 20%, cooldown remaining: none)", got)
}
