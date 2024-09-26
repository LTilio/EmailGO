package contract

type CampaignResponse struct {
	ID             string `json:"id" `
	Name           string `json:"name"`
	Content        string `json:"content"`
	Status         string `json:"status"`
	AmountOfEmails int    `json:"amount_of_emails"`
	CreatedBy      string `json:"created_by"`
}
