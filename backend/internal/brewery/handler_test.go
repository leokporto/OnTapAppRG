package brewery

import (
	"context"
	"testing"
)

type mockBeerStore struct {
	createFn  func(ctx context.Context, beer *Brewery) error
	listFn    func(ctx context.Context) ([]Brewery, error)
	getByIdFn func(ctx context.Context, beerId int64) (Brewery, error)
	findFn    func(ctx context.Context, beer *Brewery, filter string) ([]Brewery, error)
}

func (m *mockBeerStore) Create(ctx context.Context, beer *Brewery) error {
	return m.createFn(ctx, beer)
}

func (m *mockBeerStore) List(ctx context.Context) ([]Brewery, error) {
	return m.listFn(ctx)
}

func (m *mockBeerStore) GetById(ctx context.Context, beerId int64) (Brewery, error) {
	return m.getByIdFn(ctx, beerId)
}

func (m *mockBeerStore) Find(ctx context.Context, beer *Brewery, filter string) ([]Brewery, error) {
	return m.findFn(ctx, beer, filter)
}

func TestGetBreweryByID_Ok(t *testing.T) {
	mockStore := &mockBeerStore{
		getByIdFn: func(ctx context.Context, beerId int64) (Brewery, error) {
			return Brewery{ID: beerId, Name: "Test Brewery"}, nil
		},
	}

	handler := NewHandler(mockStore)

	brewery, err := getBreweryByID(context.Background(), handler.store, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if brewery.ID != 1 {
		t.Errorf("expected brewery ID to be 1, got %d", brewery.ID)
	}
}

func TestGetBreweryByID_InvalidID(t *testing.T) {
	mockStore := &mockBeerStore{}

	handler := NewHandler(mockStore)

	_, err := getBreweryByID(context.Background(), handler.store, -1)
	if err == nil {
		t.Fatalf("expected error for invalid ID, got nil")
	}
}
