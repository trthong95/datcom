package service

import (
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/people_in_charge"
	"git.d.foundation/datcom/backend/src/store/user"
)

func TestService_AddPeopleInCharge(t *testing.T) {
	type fields struct {
		User           user.ServiceMock
		Order          order.ServiceMock
		PeopleInCharge people_in_charge.ServiceMock
	}
	type args struct {
		p *domain.PeopleInCharge
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.PeopleInCharge
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when AddPeopleInCharge() success",
			fields: fields{
				user.ServiceMock{},
				order.ServiceMock{},
				people_in_charge.ServiceMock{
					ExistFunc: func(o *domain.PeopleInCharge) (bool, error) {
						return false, nil
					},
					AddFunc: func(o *domain.PeopleInCharge) (*models.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when AddPeopleInCharge() failed",
			fields: fields{
				user.ServiceMock{},
				order.ServiceMock{},
				people_in_charge.ServiceMock{
					ExistFunc: func(o *domain.PeopleInCharge) (bool, error) {
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
				User:           &tt.fields.User,
				Order:          &tt.fields.Order,
				PeopleInCharge: &tt.fields.PeopleInCharge,
			}
			got, err := s.AddPeopleInCharge(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddPeopleInCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddPeopleInCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetPeopleInCharge(t *testing.T) {
	type fields struct {
		User           user.ServiceMock
		Order          order.ServiceMock
		PeopleInCharge people_in_charge.ServiceMock
	}
	type args struct {
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.PeopleInCharge
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when GetPeopleInCharge() success",
			fields: fields{
				user.ServiceMock{},
				order.ServiceMock{},
				people_in_charge.ServiceMock{
					GetFunc: func(menuID int) ([]*domain.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when GetPeopleInCharge() failed",
			fields: fields{
				user.ServiceMock{},
				order.ServiceMock{},
				people_in_charge.ServiceMock{
					GetFunc: func(menuID int) ([]*domain.PeopleInCharge, error) {
						return nil, errors.New("Failed to get")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				User:           &tt.fields.User,
				Order:          &tt.fields.Order,
				PeopleInCharge: &tt.fields.PeopleInCharge,
			}
			got, err := s.GetPeopleInCharge(tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetPeopleInCharge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetPeopleInCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
