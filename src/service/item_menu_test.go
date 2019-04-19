package service

import (
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
)

func TestService_AddItems(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
		Item item.ServiceMock
	}
	type args struct {
		items *item.Items
		mn    *menu.Menu
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
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return &models.Menu{
							ID:       22,
							MenuName: "lunch_22_01_1995",
						}, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*item.Item, *menu.Menu) (bool, error) {
						return false, nil
					},
					AddAnItemFunc: func(*item.Item) (*models.Item, error) {
						return &models.Item{
							ItemName: "Mon 7",
							MenuID:   22,
						}, nil
					},
				},
			},
			args: args{
				items: &item.Items{
					ItemNames: []item.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				mn: &menu.Menu{
					MenuName: "lunch_22_01_1995",
				},
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
			name: "NoMenu",
			fields: fields{
				menu.ServiceMock{
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return nil, errors.New("Non-exist-menu")

					},
				},
				item.ServiceMock{},
			},
			args: args{
				mn: &menu.Menu{
					MenuName: "lunch_31_02_1995",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetMenuFailed",
			fields: fields{
				menu.ServiceMock{
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return nil, errors.New("Find Error")
					},
				},
				item.ServiceMock{},
			},
			args: args{
				mn: &menu.Menu{
					MenuName: "lunch_22_01_1995",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CheckItemExistFailed",
			fields: fields{
				menu.ServiceMock{
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return &models.Menu{
							ID:       22,
							MenuName: "lunch_22_01_1995",
						}, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*item.Item, *menu.Menu) (bool, error) {
						return false, errors.New("Check Error")
					},
				},
			},
			args: args{
				items: &item.Items{
					ItemNames: []item.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				mn: &menu.Menu{
					MenuName: "lunch_22_01_1995",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ItemExists",
			fields: fields{
				menu.ServiceMock{
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return &models.Menu{
							ID:       22,
							MenuName: "lunch_22_01_1995",
						}, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*item.Item, *menu.Menu) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				items: &item.Items{
					ItemNames: []item.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				mn: &menu.Menu{
					MenuName: "lunch_22_01_1995",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "AddItemFailed",
			fields: fields{
				menu.ServiceMock{
					FindMenuFunc: func(*menu.Menu) (*models.Menu, error) {
						return &models.Menu{
							ID:       22,
							MenuName: "lunch_22_01_1995",
						}, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(*item.Item, *menu.Menu) (bool, error) {
						return false, nil
					},
					AddAnItemFunc: func(*item.Item) (*models.Item, error) {
						return nil, errors.New("Add Item Failed")
					},
				},
			},
			args: args{
				items: &item.Items{
					ItemNames: []item.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				mn: &menu.Menu{
					MenuName: "lunch_22_01_1995",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Menu: &tt.fields.Menu,
				Item: &tt.fields.Item,
			}
			got, err := s.AddItems(tt.args.items, tt.args.mn)
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
