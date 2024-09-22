package campaign

type Repository interface {
	Create(campaign *Campaign) error
	GetBy(id string) (*Campaign, error)
	Delete(campaign *Campaign) error
	Update(campaign *Campaign) error
	GetCampaignsToBeSent() ([]Campaign, error)
}
