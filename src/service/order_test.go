package service

import (
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/order"
)

func TestService_AddOrder(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		o *order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test AddOrder() when order does not exist",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(o *order.Order) (bool, error) {
						return false, nil
					},
					AddFunc: func(o *order.Order) (*models.Order, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test AddOrder() when order exists",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(o *order.Order) (bool, error) {
						return true, nil
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Order: &tt.fields.Order,
			}
			if _, err := s.AddOrder(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Service.AddOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteOrder(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		o *order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test DeleteOrder() when order exists",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(o *order.Order) (bool, error) {
						return true, nil
					},
					DeleteFunc: func(o *order.Order) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test DeleteOrder() when order does not exist",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(o *order.Order) (bool, error) {
						return false, nil
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Order: &tt.fields.Order,
			}
			if err := s.DeleteOrder(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetOrdersByUserID(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*item.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetOrder()",
			fields: fields{
				order.ServiceMock{
					GetFunc: func(userID int) ([]*item.Item, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Order: &tt.fields.Order,
			}
			got, err := s.GetOrdersByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetOrdersByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetOrdersByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
