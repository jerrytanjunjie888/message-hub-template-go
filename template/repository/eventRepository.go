package repository

import (
	"errors"
	"fmt"

	"github.com/jerry/message-hub-template-go/template/model"
	"github.com/jinzhu/gorm"
)

type EventRepository interface {
	Create(u *model.Event) error
	Update(u *model.Event) error
	GetAll() ([]*model.Event, error)
	GetByID(id string) (*model.Event, error)
	Delete(id string) error
}

type eventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepository {
	return &eventRepo{db}
}

// get all events
func (u *eventRepo) GetAll() ([]*model.Event, error) {
	// put logs here
	events := make([]*model.Event, 0)
	err := u.db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Create a event
func (u *eventRepo) Create(event *model.Event) error {
	// put logs here
	err := u.db.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

// find a event
func (u *eventRepo) GetByID(id string) (*model.Event, error) {
	event := &model.Event{}
	err := u.db.Where("user_id = ?", id).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

//  Update event details
func (u *eventRepo) Update(event *model.Event) error {
	err := u.db.Model(&event).Update(model.Event{
		FirstName: event.FirstName,
		LastName:  event.LastName,
		UserName:  event.UserName,
		Password:  event.Password,
		Email:     event.Email}).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete event Details
func (u *eventRepo) Delete(id string) error {
	if u.db.Delete(&model.Event{}, "user_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the user with id : %s", id)
		return errors.New(errMsg)
	}
	return nil
}
