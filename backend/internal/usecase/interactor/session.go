package interactor

import "github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"

type SessionInteractor struct {
	Repositories dai.DataAccessInterfaces
}

func NewSessionInteractor(repositories dai.DataAccessInterfaces) *SessionInteractor {
	return &SessionInteractor{
		Repositories: repositories,
	}
}
