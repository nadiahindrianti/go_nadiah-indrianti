package controller

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestGetAllUsers(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
		{
			name: "berhasil",
			args: args{
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllUsers(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
		{
			name: "berhasil",
			args: args{
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
