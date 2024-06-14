package infrastructure

import (
	"SebStudy/domain/resume/commands"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CeMapper struct {
	handlers map[string]func(cloudevents.Event) (Command, error)
}

func NewCeMapper() *CeMapper {
	c := &CeMapper{}
	c.handlers = make(map[string]func(cloudevents.Event) (Command, error), 0)

	c.Register("resume.send", toSendResume)

	return c
}

func (cm *CeMapper) MapToCommand(c cloudevents.Event) (Command, error) {
	handler, err := cm.Get(c.Type())
	if err != nil {
		return nil, err
	}

	cmd, err := handler(c)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cm *CeMapper) Get(t string) (func(cloudevents.Event) (Command, error), error) {
	if h, ex := cm.handlers[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for %s type doesnt exist", t)
}

func (c *CeMapper) Register(ceType string, f func(cloudevents.Event) (Command, error)) {
	c.handlers[ceType] = f
}

func toSendResume(c cloudevents.Event) (Command, error) {
	cmd := commands.SendResume{}
	//TODO: Заполнить поля cmd из cloudevents.Event (c)

	// proto.Unmarshal(c.Data())
	return cmd, nil
}
