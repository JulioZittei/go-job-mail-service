package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestCreateANewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"john@email.com", "mary@email.com"}
	createdAt := time.Now()

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
	assert.GreaterOrEqual(t, campaign.CreatedAt, createdAt)
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content )
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}