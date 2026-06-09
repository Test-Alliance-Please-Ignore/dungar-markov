package dcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
