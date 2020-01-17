package dig

import (
	"github.com/jerry/message-hub-template-go/template/config"
	"github.com/jerry/message-hub-template-go/template/repository"
	"github.com/jerry/message-hub-template-go/template/service"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildProject() *dig.Container {
	// container.Provide()
	container.Provide(config.NewDb)
	container.Provide(config.NewConfig)
	container.Provide(repository.NewEventRepo)
	container.Provide(service.NewEventService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
