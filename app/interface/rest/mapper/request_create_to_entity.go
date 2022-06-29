package mapper

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/interface/rest/model"
)

func RequestCreateToEntity(r model.CreateShortRequestModel) entity.Link {
	return entity.Link{
		ID: r.URL,
	}
}
