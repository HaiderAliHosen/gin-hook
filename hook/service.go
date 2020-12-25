package hook

import (
	"errors"

	"github.com/google/uuid"
)

type iservice interface {
	CreateHookService(hook Hook) (Hook, error)
	GetAllHookService() ([]Hook, error)
	GetSingleHookService(hookID string) (Hook, error)
	deleteHookService(hookID string) error
	updateHookService(hookID string, hook Hook) (Hook, error)
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

//GetAllHookService --gets all
func (s *Service) GetAllHookService() ([]Hook, error) {
	getResults, err := s.rep.readAllHookRepo()
	if err != nil {
		return []Hook{}, errors.New("can't get")
	}
	return getResults, nil
}

//GetSingleHookService -- get Single data
func (s *Service) GetSingleHookService(hookID string) (Hook, error) {
	getSingleResult, err := s.rep.readSingleHookRepo(hookID)
	if err != nil {
		return Hook{}, errors.New("can't get")
	}
	return getSingleResult, nil
}

func (s *Service) deleteHookService(hookID string) error {
	err := s.rep.deleteSingleHookRepo(hookID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) updateHookService(hookID string, hook Hook) (Hook, error) {
	updateHService, err := s.rep.updateHookRepo(hookID, hook)
	if err != nil {
		return Hook{}, errors.New("cannot update")
	}
	return updateHService, nil
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
