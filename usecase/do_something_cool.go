package usecase

import (
	"context"
	"fmt"
)

type DoSomethingService struct {
	repo SuperDataStore
}

func NewDoer(repo SuperDataStore) *DoSomethingService {
	return &DoSomethingService{
		repo: repo,
	}
}

func (s *DoSomethingService) DoTheThingThatRequiresMultiple(ctx context.Context) error {
	err := s.repo.Atomic(ctx, func(ctx context.Context, ds SuperDataStore) error {
		// now my sub repos will be wrapped in the Tx
		subRepo := ds.SomeSubDataStore()
		subRepo2 := ds.SomeOtherSubDataStore()

		t, err := subRepo.GetAThing()
		if err != nil {
			return err
		}

		s, err := subRepo2.GetSomething()
		if err != nil {
			return err
		}

		fmt.Println(t)
		fmt.Println(s)

		return subRepo2.SaveSomething(s)
	})

	return err
}
