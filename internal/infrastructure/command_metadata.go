package infrastructure

type CommandMetadata struct {
	AggregateId string
	UserId      string
}

func NewCommandMetadata(AggregateId string) CommandMetadata {
	return CommandMetadata{
		AggregateId: AggregateId,
	}
}

func (md *CommandMetadata) SetUser(userId string) {
	md.UserId = userId
}
