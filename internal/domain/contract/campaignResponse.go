package contract

type CampaignOutput struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	Status       string `json:"status"`
	EmailsToSend int    `json:"emailsToSend"`
	CreatedBy    string `json:"createdBy"`
}
