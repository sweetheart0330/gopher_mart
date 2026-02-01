package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sweetheart0330/gopher_mart/internal/mocks"
	"github.com/sweetheart0330/gopher_mart/internal/models"
)

func TestGetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUseCase := mocks.NewMockController(ctrl)

	f := func(ctx context.Context, userID string, expBalance []byte, expErr error) {
		balance := models.Balance{}
		err := json.Unmarshal(expBalance, &balance)
		if err != nil {
		}
		mockUseCase.EXPECT().GetBalance(ctx, userID).Return(balance, expErr)
	}
	ctxUser := "testUser"
	tests := []struct {
		name       string
		wantStatus int
		expError   error
		wantBody   string
		f          func(ctx context.Context, userID string, expBalance []byte, expErr error)
	}{
		{
			name:       "success",
			wantStatus: 200,
			expError:   nil,
			f:          f,
			wantBody:   "{\"id\":12}",
		},

		{
			name:       "success",
			wantStatus: 500,
			wantBody:   "some error",
			expError:   errors.New("some error"),
			f:          f,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				controller: mockUseCase, // подгони под реальные поля твоего Handler
			}

			req := httptest.NewRequest(http.MethodGet, "/balance", nil)
			ctx := context.WithValue(req.Context(), models.UserIDKey, ctxUser)
			req = req.WithContext(ctx)
			tt.f(ctx, ctxUser, []byte(tt.wantBody), tt.expError)
			rec := httptest.NewRecorder()

			h.GetBalance(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status: want %d, got %d, body=%q", tt.wantStatus, rec.Code, rec.Body.String())
				return
			}

		})
	}
}
