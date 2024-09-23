package campaign

import (
	internalerror "EmailGO/internal/internalError"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Started  string = "Started"
	Done     string = "Done"
	Fail     string = "Fail"
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
)

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;not null"`
	Name      string    `validate:"min=5,max=50" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive"` //dive para validar um struc dentro de outra
	Status    string    `gorm:"size:20;not null"`
	CreatedBy string    `validate:"required,email" gorm:"size:50;not null"`
}

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:50"`
	CampaignId string `gorm:"size:50"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = Deleted
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Fail() {
	c.Status = Fail
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Started() {
	c.Status = Started
	c.UpdatedOn = time.Now()
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
		contacts[index].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
	}
	err := internalerror.ValidateStruc(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err

}
