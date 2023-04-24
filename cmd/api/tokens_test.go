package main

import (
	"encoding/json"
	"greenlight.bcc/internal/assert"
	"net/http"
	"testing"
)

func TestCreateToken(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	tests := []struct {
		name     string
		Email    string
		Password string
		wantCode int
	}{
		{
			name:     "Create token",
			Email:    "admin",
			Password: "Qwerty1!",
			wantCode: http.StatusCreated,
		},
		{
			name:     "Unauthorized",
			Email:    "dinar@gmail.com",
			Password: "12345678",
			wantCode: http.StatusUnauthorized,
		},
		{
			name:     "Validation Fail",
			Email:    "1234",
			Password: "123",
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}{
				Email:    tt.Email,
				Password: tt.Password,
			}

			b, err := json.Marshal(&inputData)
			if err != nil {
				t.Fatal("wrong input data")
			}
			if tt.name == "test for wrong input" {
				b = append(b, 'a')
			}

			code, _, _ := ts.postForm(t, "/v1/tokens/authentication", b)

			assert.Equal(t, code, tt.wantCode)
		})
	}

	code, _, _ := ts.postForm(t, "/v1/tokens/authentication", []byte{})

	assert.Equal(t, code, http.StatusBadRequest)
}
