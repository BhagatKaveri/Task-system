package service

import "GST/billlingSystem/db"

type Dependencies struct {
	BillerService Service
}

func InitDependencies() (deps Dependencies, err error) {
	store, err := db.Init()
	if err != nil {
		return
	}

	UserService := NewBillerService(store)

	deps = Dependencies{
		BillerService: UserService,
	}
	return deps, nil

}
