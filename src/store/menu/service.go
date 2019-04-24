package menu

// Service ..
type Service interface {
	CheckMenuExist(menuID int) (bool, error)
}
