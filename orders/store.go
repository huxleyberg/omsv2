package main

import "context"

type store struct {
	//add db connection here
}

func NewStore() *store {
	return &store{}
}

func (s store) Create(ctx context.Context) error {
	//add db logic here
	return nil
}
