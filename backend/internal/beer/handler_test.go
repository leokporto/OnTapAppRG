package beer

import (
	"context"
	"testing"
)

type mockBeerStore struct {
	createFn       func(ctx context.Context, beer *Beer) error
	listAllFn      func(ctx context.Context) ([]BeerResponse, error)
	getByIDFn      func(ctx context.Context, beerId int64) (BeerResponse, error)
	getStylesFn    func(ctx context.Context) ([]BeerStyle, error)
	getBreweriesFn func(ctx context.Context) ([]Brewery, error)
}

func (m *mockBeerStore) Create(ctx context.Context, beer *Beer) error {
	return m.createFn(ctx, beer)
}

func (m *mockBeerStore) ListAll(ctx context.Context) ([]BeerResponse, error) {
	return m.listAllFn(ctx)
}

func (m *mockBeerStore) GetByID(ctx context.Context, beerId int64) (BeerResponse, error) {
	return m.getByIDFn(ctx, beerId)
}

func (m *mockBeerStore) GetStyles(ctx context.Context) ([]BeerStyle, error) {
	return m.getStylesFn(ctx)
}

func (m *mockBeerStore) GetBreweries(ctx context.Context) ([]Brewery, error) {
	return m.getBreweriesFn(ctx)
}

func TestGetBeerByID_Ok(t *testing.T) {
	mockStore := &mockBeerStore{
		getByIDFn: func(ctx context.Context, beerId int64) (BeerResponse, error) {
			return BeerResponse{ID: beerId, Name: "Test Beer", Style: "American IPA", Brewery: "Test Beer Company", FullName: "Test Beer IPA", ABV: 6.5, MinIBU: 45, MaxIBU: 55}, nil
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
