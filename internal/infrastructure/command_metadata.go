package infrastructure

type CommandMetadata struct {
	AggregateId string
	UserId      string
}

func NewCommandMetadata(AggregateId string, userId string) CommandMetadata {
	return CommandMetadata{
		AggregateId: AggregateId,
		UserId:      userId,
	}
}

func (md *CommandMetadata) SetUser(userId string) {
	md.UserId = userId
}
