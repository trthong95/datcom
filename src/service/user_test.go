package service

import (
	"datcom/backend/models"
	"datcom/backend/src/domain"
	"datcom/backend/src/store/user"
	"errors"
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
)

func TestService_GetAllUser(t *testing.T) {
	type fields struct {
		User user.ServiceMock
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				user.ServiceMock{
					FindAllFunc: func() ([]*models.User, error) {
						u := []*models.User{
							&models.User{
								Email: "demo",
								ID:    100,
								Name:  "demo",
							},
						}
						return u, nil
					},
				},
			},
			want: []*models.User{
				&models.User{
					Email: "demo",
					ID:    100,
					Name:  "demo",
				},
			},
			wantErr: false,
		},
		{
			name: "Fail",
			fields: fields{
				user.ServiceMock{
					FindAllFunc: func() ([]*models.User, error) {
						u := []*models.User{
							&models.User{
								Email: "",
								ID:    0,
								Name:  "",
							},
						}
						return u, nil
					},
				},
			},
			want: []*models.User{
				&models.User{
					Email: "",
					ID:    0,
					Name:  "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				User: &tt.fields.User,
			}
			got, err := s.GetAllUser()
			pp.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	type fields struct {
		User user.ServiceMock
	}
	type args struct {
		p *domain.PersonInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				user.ServiceMock{
					CreateFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return &models.User{
							ID:    100,
							Name:  "Demo1",
							Email: "Demo@email.com",
							Token: "ABCD",
						}, nil
					},
					FindFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return nil, nil
					},
				},
			},
			args: args{
				&domain.PersonInfo{
					Name:  "Demo1",
					Email: "Demo@email.com",
					Token: "ABCD",
				},
			},
			want: &models.User{
				ID:    100,
				Email: "demo",
				Name:  "demo",
				Token: "ABCD",
			},
			wantErr: false,
		},
		{
			name: "Duplicate",
			fields: fields{
				user.ServiceMock{
					CreateFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return &models.User{}, nil
					},
					FindFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return &models.User{
							ID:    100,
							Name:  "Demo2",
							Email: "Demo@email.com",
							Token: "ABCD",
						}, nil
					},
				},
			},
			args: args{
				&domain.PersonInfo{
					Name:  "Demo2",
					Email: "Demo@email.com",
					Token: "ABCD",
				},
			},
			want:    &models.User{},
			wantErr: true,
		},
		{
			name: "Find Error",
			fields: fields{
				user.ServiceMock{
					CreateFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return &models.User{}, errors.New("Find errors")
					},
					FindFunc: func(p *domain.PersonInfo) (*models.User, error) {
						return &models.User{}, errors.New("Find errors")
					},
				},
			},
			args: args{
				&domain.PersonInfo{},
			},
			want:    &models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				User: &tt.fields.User,
			}
			got, err := s.CreateUser(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("Service.CreateUser() = %v, want %v", got.ID, tt.want.ID)
			}
		})
	}
}
