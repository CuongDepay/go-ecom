package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CuongDepay/go-ecom/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("register user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password",
		}

		jsonPayload, _ := json.Marshal(payload)

		request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayload))

		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status %d but got %d", http.StatusBadRequest, recorder.Code)
		}

	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
