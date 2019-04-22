package service

import (
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/pic"
)

func TestService_AddPIC(t *testing.T) {
	type fields struct {
		PIC pic.ServiceMock
	}
	type args struct {
		p *domain.PICInput
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
			name: "when AddPIC() success",
			fields: fields{
				pic.ServiceMock{
					ExistFunc: func(o *domain.PICInput) (bool, error) {
						return false, nil
					},
					AddFunc: func(o *domain.PICInput) (*models.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when AddPIC() failed",
			fields: fields{
				pic.ServiceMock{
					ExistFunc: func(o *domain.PICInput) (bool, error) {
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
				Store: Store{
					PIC: &tt.fields.PIC,
				},
			}
			got, err := s.AddPIC(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddPIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddPIC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetPICByMenuID(t *testing.T) {
	type fields struct {
		PIC pic.ServiceMock
	}
	type args struct {
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.PeopleInCharge
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when GetPIC() success",
			fields: fields{
				pic.ServiceMock{
					GetByMenuIDFunc: func(menuID int) ([]*models.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when GetPIC() failed",
			fields: fields{
				pic.ServiceMock{
					GetByMenuIDFunc: func(menuID int) ([]*models.PeopleInCharge, error) {
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
				Store: Store{
					PIC: &tt.fields.PIC,
				},
			}
			got, err := s.GetPICByMenuID(tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetPICByMenuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetPICByMenuID() = %v, want %v", got, tt.want)
			}
		})
	}
}
