package hook

import (
	"errors"

	"github.com/google/uuid"
)

type iservice interface {
	CreateHookService(hook Hook) (Hook, error)
}

//Service struct
type Service struct {
	rep *repo
}

//CreateHookService ---func
func (s *Service) CreateHookService(hook Hook) (Hook, error) {
	hook.ID = generateUUID()
	postResults, err := s.rep.createHookRepo(hook)
	if err != nil {
		return Hook{}, errors.New("cannot create")
	}
	return postResults, nil
}

func generateUUID() string {
	return uuid.New().String()
}

//NewHookService --- func
func NewHookService(r *repo) *Service {
	return &Service{
		rep: r,
	}
}
