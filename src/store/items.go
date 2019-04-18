package store

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
)

func (store *storePostgres) InsertItemsToDB(itemNames []string, menuName string) error {
	var err error
	tx, err := store.db.BeginTx(context.Background(), nil)
	defer func() {
		if err == nil {
			tx.Commit()
			return
		}
		pp.Println("An error occured, rollback...")
		tx.Rollback()
	}()
	menu, findErr := models.Menus(qm.Where("menu_name = ?", menuName)).One(context.Background(), tx)
	pp.Println(menu)
	if findErr != nil {
		err = findErr
		pp.Println("Failed to get menu id...")
	} else {
		for _, name := range itemNames {
			// Check if item is duplicated
			exists, ExErr := models.Items(
				qm.Where("item_name = ?", name),
			).Exists(context.Background(), tx)
			if ExErr != nil {
				err = ExErr
				pp.Println("Failed to check if item exists...")
			}
			if !exists {
				// Not duplicate, insert it
				item := &models.Item{ItemName: name, MenuID: menu.ID}
				InsertErr := item.Insert(context.Background(), tx, boil.Infer())
				if InsertErr != nil {
					err = InsertErr
					pp.Println("Failed to insert item into DB...")
				}
			} else {
				// Duplicate, skip it
				pp.Println("Item already exist, skipping...")
			}
			if err != nil {
				break
			}
		}
	}
	return err
}
