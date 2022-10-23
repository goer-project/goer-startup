package v1

import "goer-startup/internal/apiserver/store"

// Service defines functions used to return resource interface.
type Service interface {
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
