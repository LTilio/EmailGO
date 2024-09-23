package campaign

import (
	"EmailGO/internal/contract"
	internalerror "EmailGO/internal/internalError"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaignRequest) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Delete(id string) error
	Start(id string) error
}

type ServiceImp struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaignRequest) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	if err != nil {
		return "", err
	}
	err = s.Repository.Create(campaign)
	if err != nil {
		return "", internalerror.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {

	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerror.ProcessError(err)
	}

	return &contract.CampaignResponse{
		ID:             campaign.ID,
		Name:           campaign.Name,
		Content:        campaign.Content,
		Status:         campaign.Status,
		AmountOfEmails: len(campaign.Contacts),
		CreatedBy:      campaign.CreatedBy,
	}, nil
}

func (s *ServiceImp) Delete(id string) error {

	campaign, err := s.getAndValidateStatusIsPending(id)

	if err != nil {
		return err
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerror.ErrInternal
	}
	return nil
}

func (s *ServiceImp) SendMailAndUpdateStatus(campaignSaved *Campaign) {

	err := s.SendMail(campaignSaved)
	if err != nil {
		campaignSaved.Fail()
	} else {
		campaignSaved.Done()
	}
	s.Repository.Update(campaignSaved)

}

// TODO: make unit test
func (s *ServiceImp) Start(id string) error {

	campaignSaved, err := s.getAndValidateStatusIsPending(id)

	if err != nil {
		return err
	}

	// dá pra chamar metodos também, só colocar o go na frente e esse metodo vai ser chamado em paralero
	// go s.SendMailAndUpdateStatus(campaignSaved)

	// //função anonima em paralelo - melhorando a performance do metodo
	// go func() {
	// 	err := s.SendMail(campaignSaved)
	// 	if err != nil {
	// 		campaignSaved.Fail()
	// 	} else {
	// 		campaignSaved.Done()
	// 	}
	// 	s.Repository.Update(campaignSaved)
	// }()

	campaignSaved.Started()
	err = s.Repository.Update(campaignSaved)
	if err != nil {
		return internalerror.ErrInternal
	}

	return nil
}

func (s *ServiceImp) getAndValidateStatusIsPending(id string) (*Campaign, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerror.ProcessError(err)
	}

	if campaign.Status != Pending {
		return nil, errors.New("Campaign status invalid")
	}
	return campaign, nil
}
