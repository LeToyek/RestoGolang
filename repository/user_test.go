package repository

import (
	"regexp"
	"resto/entities"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepo_AddUser(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}

	user := entities.User{
		First_name: "toyek",
		Last_name:  "goblog",
		Email:      "toyek@goblog.com",
		Password:   "toyek",
		Phone:      "12131213",
		Created_at: time.Now().String(),
		Updated_at: time.Now().String(),
		User_id:    "41213",
	}
	tests := []struct {
		name    string
		param   entities.User
		wantErr error
		mock    func(mock mockFields)
	}{
		{
			name:    "case error",
			param:   user,
			wantErr: assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectExec(regexp.QuoteMeta(queryInsertUser)).
					WithArgs(
						user.First_name,
						user.Last_name,
						user.Email,
						user.Password,
						user.Phone,
						user.Created_at,
						user.Updated_at,
						user.User_id,
					).WillReturnError(assert.AnError)
			},
		},
		{
			name:    "case success",
			param:   user,
			wantErr: nil,
			mock: func(mock mockFields) {
				mock.sql.ExpectExec(regexp.QuoteMeta(queryInsertUser)).
					WithArgs(
						user.First_name,
						user.Last_name,
						user.Email,
						user.Password,
						user.Phone,
						user.Created_at,
						user.Updated_at,
						user.User_id,
					).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			assert.Nil(t, err)
			defer mockDB.Close()

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)

			repository := &Repository{
				DB: mockDB,
			}

			gotErr := repository.AddUser(test.param)
			assert.Equal(t, test.wantErr, gotErr)
		})
	}
}

func TestRepo_GetUsers(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}
	user := entities.User{
		First_name: "toyek",
		Last_name:  "goblog",
		Email:      "toyek@goblog.com",
		Password:   "toyek",
		Phone:      "12131213",
		Created_at: time.Now().String(),
		Updated_at: time.Now().String(),
		User_id:    "41213",
	}
	tests := []struct {
		name     string
		wantUser entities.User
		wantErr  error
		mock     func(mock mockFields)
	}{
		{
			name:     "case query error",
			wantUser: user,
			wantErr:  assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectQuery(regexp.QuoteMeta(queryGetAllUsers)).
					WillReturnRows(sqlmock.NewRows([]string{
						"first_name",
						"last_name",
						"email",
						"password",
						"phone",
						"created_at",
						"updated_at",
						"user_id"}).AddRow(
						user.First_name,
						user.Last_name,
						user.Email,
						user.Password,
						user.Phone,
						user.Created_at,
						user.Updated_at,
						user.User_id))
				// mock.sql.ExpectQuery(regexp.QuoteMeta(queryGetAll)).
				// 	WillReturnError()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			assert.Nil(t, err)
			defer mockDB.Close()

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)

			repository := &Repository{
				DB: mockDB,
			}

			users, err := repository.GetUsers()
			assert.Equal(t, test.wantUser, users)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
