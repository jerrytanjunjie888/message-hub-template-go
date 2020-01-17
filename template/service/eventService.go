package service

import (
	"github.com/jerry/message-hub-template-go/template/model"
	"github.com/jerry/message-hub-template-go/template/repository"
)

type Service interface {
	Create(u *model.Event) error
	Update(u *model.Event) error
	GetByID(id string) (*model.Event, error)
	GetAll() ([]*model.Event, error)
	Delete(id string) error
}

// init eventRepository
type eventService struct {
	repo repository.EventRepository
}

// realize eventRepository
func NewEventService(repo repository.EventRepository) Service {
	return &eventService{
		repo: repo,
	}
}

func (svc *eventService) Create(e *model.Event) error { return svc.repo.Create(e) }

func (svc *eventService) GetAll() ([]*model.Event, error) { return svc.repo.GetAll() }

func (svc *eventService) GetByID(id string) (*model.Event, error) { return svc.repo.GetByID(id) }

func (svc *eventService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *eventService) Update(u *model.Event) error { return svc.repo.Update(u) }
