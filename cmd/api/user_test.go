package main

import (
	"encoding/json"
	"greenlight.bcc/internal/assert"
	"net/http"
	"testing"
)

//registerUserHandler
//activateUserHandler

func TestRegisterUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	tests := []struct {
		name     string
		Name     string
		Email    string
		Password string
		wantCode int
	}{
		{
			name:     "Register user",
			Name:     "User",
			Email:    "user@gmail.com",
			Password: "Qwerty1!",
			wantCode: http.StatusCreated,
		},
		{
			name:     "Duplicate email",
			Name:     "Admin",
			Email:    "admin@gmail.com",
			Password: "Qwerty1!",
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "Fail validation",
			Name:     "",
			Email:    "",
			Password: "",
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := struct {
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				Name:     tt.Name,
				Email:    tt.Email,
				Password: tt.Password,
			}

			body, err := json.Marshal(&input)
			if err != nil {
				t.Fatal("wrong input data")
			}
			if tt.name == "test for wrong input" {
				body = append(body, 'a')
			}

			code, _, _ := ts.postForm(t, "/v1/users", body)

			assert.Equal(t, code, tt.wantCode)

		})
	}
}

func TestActivateUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	tests := []struct {
		name           string
		TokenPlainText string
		wantCode       int
	}{
		{
			name:           "Activate user",
			TokenPlainText: "TokenPlainTextForTokenTest",
			wantCode:       http.StatusOK,
		},
		{
			name:           "Fail validation",
			TokenPlainText: "",
			wantCode:       http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := struct {
				TokenPlainText string `json:"token"`
			}{
				TokenPlainText: tt.TokenPlainText,
			}

			b, err := json.Marshal(&inputData)
			if err != nil {
				t.Fatal("wrong input data")
			}
			if tt.name == "test for wrong input" {
				b = append(b, 'a')
			}

			code, _, _ := ts.putForm(t, "/v1/users/activated", b)

			assert.Equal(t, code, tt.wantCode)

		})
	}
}
