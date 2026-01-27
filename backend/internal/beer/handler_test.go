package beer

import (
	"context"
	"testing"
)

type mockBeerStore struct {
	createFn        func(ctx context.Context, beer *Beer) error
	listFn          func(ctx context.Context) ([]Beer, error)
	getByIdFn       func(ctx context.Context, beerId int64) (Beer, error)
	findFn          func(ctx context.Context, beer *Beer, filter string) ([]Beer, error)
	listByBreweryFn func(ctx context.Context, breweryId int64) ([]Beer, error)
}

func (m *mockBeerStore) Create(ctx context.Context, beer *Beer) error {
	return m.createFn(ctx, beer)
}

func (m *mockBeerStore) List(ctx context.Context) ([]Beer, error) {
	return m.listFn(ctx)
}

func (m *mockBeerStore) GetById(ctx context.Context, beerId int64) (Beer, error) {
	return m.getByIdFn(ctx, beerId)
}

func (m *mockBeerStore) Find(ctx context.Context, beer *Beer, filter string) ([]Beer, error) {
	return m.findFn(ctx, beer, filter)
}

func (m *mockBeerStore) ListByBrewery(ctx context.Context, breweryId int64) ([]Beer, error) {
	return m.listByBreweryFn(ctx, breweryId)
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
