package user

import (
	"errors"
	"timtube/api"
	"timtube/domain"
)

const ChannelSubscriptionURL = api.BASE_URL + "channel/subscription/"

func CreateChannel(use domain.ChannelSubscription) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().SetBody(use).Post(ChannelSubscriptionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateChannel(role domain.ChannelSubscription) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().SetBody(role).Post(ChannelSubscriptionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadChannel(id string) (domain.ChannelSubscription, error) {
	entity := domain.ChannelSubscription{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannelByUserId(id string) ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get-by-user/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadChannelByChannel(id string) ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "get-by-channel/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteChannel(id string) (domain.Channel, error) {
	entity := domain.Channel{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadChannels() ([]domain.Channel, error) {
	entity := []domain.Channel{}
	resp, _ := api.Rest().Get(ChannelSubscriptionURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
