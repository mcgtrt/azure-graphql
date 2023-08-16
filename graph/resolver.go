package graph

import "github.com/mcgtrt/azure-graphql/store"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store store.EmployeeStorer
}

func NewResolver() (*Resolver, error) {
	store, err := store.NewAzureEmployeeStore()
	if err != nil {
		return nil, err
	}
	return &Resolver{
		store: store,
	}, nil
}
