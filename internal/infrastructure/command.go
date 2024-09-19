package infrastructure

type Command interface {
	GetAggregateId() string
}

type BaseCommand struct {
	AggregateId string
}

func NewBaseCommand(aggregateId string) BaseCommand {
	return BaseCommand{AggregateId: aggregateId}
}

func (c *BaseCommand) GetAggregateId() string {
	return c.AggregateId
}
