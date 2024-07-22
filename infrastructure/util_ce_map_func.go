package infrastructure

import (
	"encoding/base64"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

func DecodeCloudeventData(c cloudevents.Event, protoMes proto.Message) error {

	if err := c.DataAs(protoMes); err == nil {
		return nil
	}

	err := proto.Unmarshal(c.Data(), protoMes)
	if err == nil {
		return nil
	}

	bytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(bytes, protoMes); err == nil {
		return nil
	}

	return fmt.Errorf("failed to decode cloudevent data")
}
