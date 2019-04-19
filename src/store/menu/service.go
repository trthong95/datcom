package menu

import "git.d.foundation/datcom/backend/models"

// Service ..
type Service interface {
	FindMenu(m *Menu) (*models.Menu, error)
}
