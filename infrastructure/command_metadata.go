package infrastructure

type CommandMetadata struct {
	AggregateId string
	From        string
}

func NewCommandMetadata(AggregateId string) CommandMetadata {
	return CommandMetadata{
		AggregateId: AggregateId,
	}
}
