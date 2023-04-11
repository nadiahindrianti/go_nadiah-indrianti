package controllers

import (
	"net/http"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func TestGetUsersController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUsersController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUsersController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "kartika123",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBooksController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				Judul:    "mentari",
				Penulis:  "kartika",
				Penerbit: "PT.KaryaSastra",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				Judul:    "nil",
				Penulis:  "nil",
				Penerbit: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBooksController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBooksController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				Judul:    "mentari",
				Penulis:  "kartika",
				Penerbit: "PT.KaryaSastra",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				Judul:    "nil",
				Penulis:  "nil",
				Penerbit: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				Judul:    "nil",
				Penulis:  "nil",
				Penerbit: "nil",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				Judul:    "mentari",
				Penulis:  "kartika",
				Penerbit: "PT.KaryaSastra",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				Judul:    "mentari",
				Penulis:  "kartikasari",
				Penerbit: "PT.KaryaSastra",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				Judul:    "mentari",
				Penulis:  "kartika",
				Penerbit: "PT.KaryaSastra",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUserControllers(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoginUserControllers(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LoginUserControllers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserDetailController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid test",
			args: args{
				name:     "kartika",
				Email:    "kartika123@gmail.com",
				Password: "123katika",
			},
			wantErr: http.StatusOK,
		},

		{
			name: "invalid test",
			args: args{
				name:     "nil",
				Email:    "nil",
				Password: "nil",
			},
			wantErr: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUserDetailController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUserDetailController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
