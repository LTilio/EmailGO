package database

import (
	"EmailGO/internal/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)
	return tx.Error
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := c.Db.Preload("Contacts").First(&campaign, "id=?", id)
	return &campaign, tx.Error
}

// metodo de delete com RollBack
func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {

	tx := c.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Select("Contacts").Delete(campaign).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
