package openapi

import (
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
)

func apiContact(v app.Contact) *model.Contact {
	return &model.Contact{
		ID:   int32(v.ID),
		Name: &v.Name,
	}
}

func apiContacts(vs []app.Contact) []*model.Contact {
	ms := make([]*model.Contact, len(vs))
	for i := range vs {
		ms[i] = apiContact(vs[i])
	}
	return ms
}
