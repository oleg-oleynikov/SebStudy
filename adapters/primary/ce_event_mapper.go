package primary

import (
	"SebStudy/domain/resume/commands"
	"fmt"
	"reflect"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

type CeMapper struct {
	handlers map[string]func(cloudevents.Event) (interface{}, error)
}

func NewCeMapper() *CeMapper {
	c := &CeMapper{}
	c.handlers = make(map[string]func(cloudevents.Event) (interface{}, error), 0)

	// Регистрация handler а, для преобразования cloudevents.Event в объект понятный приложению
	c.Register("resume.send", toSendResume)

	return c
}

func (cm *CeMapper) MapToCommand(c cloudevents.Event) (interface{}, error) {
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

func (cm *CeMapper) Get(t string) (func(cloudevents.Event) (interface{}, error), error) {
	if h, ex := cm.handlers[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for %s type doesnt exist", t)
}

func (c *CeMapper) Register(ceType string, f func(cloudevents.Event) (interface{}, error)) {
	c.handlers[ceType] = f
}

func toSendResume(c cloudevents.Event) (interface{}, error) {
	cmd := commands.SendResume{}
	var de pb.ResumeSended
	proto.Unmarshal(c.Data(), &de)
	fmt.Printf("Event data: %v\nEvent type: %s\n", &de, reflect.ValueOf(&de))
	//TODO: Заполнить поля cmd из cloudevents.Event (c)
	// proto.Unmarshal(c.Data())
	return cmd, nil
}
