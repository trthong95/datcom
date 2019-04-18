package user

import (
	"datcom/backend/models"
	"datcom/backend/src/domain"
)

type Service interface {
	Create(p *domain.PersonInfo) (*models.User, error)
	Find(p *domain.PersonInfo) (*models.User, error)
	FindAll() ([]*models.User, error)
}
