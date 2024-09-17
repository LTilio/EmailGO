package campaign

type Repository interface {
	Create(campaign *Campaign) error
	GetBy(id string) (*Campaign, error)
	Delete(campaign *Campaign) error
}
