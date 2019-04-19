package service

import (
	"database/sql"
	"testing"
)

func TestNewService(t *testing.T) {
	mydb := &sql.DB{}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "coverage",
			args: args{
				mydb,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.db); got == nil {
				t.Errorf("NewService() = %v, want not nil", got)
			}
		})
	}
}
