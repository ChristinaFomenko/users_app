package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ChristinaFomenko/users_app/pkg/database"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"

	mock_database "github.com/ChristinaFomenko/users_app/pkg/database/mocks"
	"github.com/ChristinaFomenko/users_app/pkg/model"
)

func TestCreateUserHandler(t *testing.T) {
	type mockBehavior func(s *mock_database.MockUserDB, user model.User)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstname":"Test","lastname":"Test","dateofbirth":"2000","incomeperyear":"10.99"}`,
			inputUser: model.User{
				FirstName:     "Christina",
				LastName:      "Fomenko",
				DateOfBirth:   2000,
				IncomePerYear: 10.99,
			},
			mockBehavior: func(s *mock_database.MockUserDB, user model.User) {
				s.EXPECT().CreateUser(user).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			create := mock_database.NewMockUserDB(c)
			testCase.mockBehavior(create, testCase.inputUser)

			repo := &database.UserRepository{}
			handler := NewApp(nil, repo)

			r := mux.NewRouter()
			r.HandleFunc("/user/create", CreateUserHandler(handler.repoUser))

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/create", bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
