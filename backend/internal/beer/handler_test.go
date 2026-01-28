package beer

import (
	"context"
	"testing"
)

type mockBeerStore struct {
	createFn  func(ctx context.Context, beer *Beer) error
	getByIdFn func(ctx context.Context, beerId int64) (Beer, error)
}

func (m *mockBeerStore) Create(ctx context.Context, beer *Beer) error {
	return m.createFn(ctx, beer)
}

func (m *mockBeerStore) GetById(ctx context.Context, beerId int64) (Beer, error) {
	return m.getByIdFn(ctx, beerId)
}

func TestGetBeerByID_Ok(t *testing.T) {
	mockStore := &mockBeerStore{
		getByIdFn: func(ctx context.Context, beerId int64) (Beer, error) {
			return Beer{ID: beerId, Name: "Test Beer", StyleID: 5, BreweryID: 3, FullName: "Test Beer IPA", ABV: 6.5, MinIBU: 45, MaxIBU: 55}, nil
		},
	}

	handler := NewHandler(mockStore)

	beer, err := getBeerByID(context.Background(), handler.store, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if beer.ID != 1 {
		t.Errorf("expected beer ID to be 1, got %d", beer.ID)
	}
}

func TestGetBeerByID_InvalidID(t *testing.T) {
	mockStore := &mockBeerStore{}

	handler := NewHandler(mockStore)

	_, err := getBeerByID(context.Background(), handler.store, -1)
	if err == nil {
		t.Fatalf("expected error for invalid ID, got nil")
	}
}
