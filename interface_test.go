package notify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	//assert.Implements(t, (*notifier)(nil), new(Email))
	//assert.Implements(t, (*notifier)(nil), new(Slack))
	//assert.Implements(t, (*notifier)(nil), new(Telegram))
	assert.Implements(t, (*notifier)(nil), new(Webhook))
}