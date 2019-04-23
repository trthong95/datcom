package service

import (
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
)

func TestService_AddItems(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
		Item item.ServiceMock
	}
	type args struct {
		items  *domain.ItemInput
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*domain.Item) (bool, error) {
						return false, nil
					},
					AddFunc: func(*domain.Item) (*models.Item, error) {
						return &models.Item{
							ItemName: "Mon 7",
							MenuID:   22,
						}, nil
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want: []*models.Item{
				&models.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			wantErr: false,
		},
		{
			name: "CheckMenuExistFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return false, errors.New("Find Error")
					},
				},
				item.ServiceMock{},
			},
			args: args{
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MenuNotExist",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return false, nil
					},
				},
				item.ServiceMock{},
			},
			args: args{
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CheckItemExistFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*domain.Item) (bool, error) {
						return false, errors.New("Check Error")
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ItemExists",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*domain.Item) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "AddItemFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*domain.Item) (bool, error) {
						return false, nil
					},
					AddFunc: func(*domain.Item) (*models.Item, error) {
						return nil, errors.New("Add Item Failed")
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store{
					Menu: &tt.fields.Menu,
					Item: &tt.fields.Item,
				},
			}
			got, err := s.AddItems(tt.args.items, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
