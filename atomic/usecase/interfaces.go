package usecase

import "context"

type SuperDataStore interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, ds SuperDataStore) error) error
	SomeSubDataStore() SomeSubDataStore
	SomeOtherSubDataStore() SomeOtherSubDataStore
}

type SomeSubDataStore interface {
	GetAThing() (*AThing, error)
	SaveAThing(thing *AThing) error
}

type SomeOtherSubDataStore interface {
	GetSomething() (*Something, error)
	SaveSomething(s *Something) error
}
