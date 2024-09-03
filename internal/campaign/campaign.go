package campaign

import (
	internalerror "EmailGO/internal/internalError"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=50"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"` //dive para validar um struc dentro de outra
}

type Contact struct {
	Email string `validate:"email"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	constacts := make([]Contact, len(emails))
	for index, value := range emails {
		constacts[index].Email = value
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  constacts,
	}
	err := internalerror.ValidateStruc(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err

}
