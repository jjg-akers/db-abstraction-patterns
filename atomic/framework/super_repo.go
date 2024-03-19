package framework

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"

	"github.com/jjg-akers/db-abstraction-patterns/atomic/usecase"
)

type AtomicRepo struct {
	be   beginner
	exec contextExecuter
}

func NewAtomicRepo(db beginner) *AtomicRepo {
	// On initialization, set both to the beginner
	return &AtomicRepo{
		be:   db,
		exec: db,
	}
}

func (r *AtomicRepo) Atomic(ctx context.Context, fn func(ctx context.Context, repo usecase.SuperDataStore) error) (err error) {
	tx, err := r.be.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()

			switch e := p.(type) {
			case runtime.Error:
				panic(e)
			case error:
				err = fmt.Errorf("panic err: %v", p)
				return
			default:
				panic(e)
			}
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	// inject new tx as the executer
	txRepo := &AtomicRepo{
		be:   r.be,
		exec: tx,
	}

	// call the fn with our new tx executer
	err = fn(ctx, txRepo)
	return
}

func (r *AtomicRepo) SomeSubDataStore() usecase.SomeSubDataStore {
	repo, _ := NewSomeSubDataStore(r.exec)
	return repo
}
func (r *AtomicRepo) SomeOtherSubDataStore() usecase.SomeOtherSubDataStore {
	repo, _ := NewSomeOtherSubDataStore(r.exec)
	return repo
}

// Sub Repos
type SomeSubDataStore struct {
	db contextExecuter
}

func (r *SomeSubDataStore) GetAThing() (*usecase.AThing, error) {
	return nil, nil
}
func (r *SomeSubDataStore) SaveAThing(thing *usecase.AThing) error {
	return nil
}

func NewSomeSubDataStore(db contextExecuter) (*SomeSubDataStore, error) {
	return &SomeSubDataStore{
		db: db,
	}, nil
}

type SomeOtherSubDataStore struct {
	db contextExecuter
}

func (r *SomeOtherSubDataStore) GetSomething() (*usecase.Something, error) {
	return nil, nil
}
func (r *SomeOtherSubDataStore) SaveSomething(s *usecase.Something) error {
	return nil
}

func NewSomeOtherSubDataStore(db contextExecuter) (*SomeOtherSubDataStore, error) {
	return &SomeOtherSubDataStore{
		db: db,
	}, nil
}
