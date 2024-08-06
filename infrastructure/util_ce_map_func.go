package infrastructure

import (
	"google.golang.org/protobuf/proto"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

func DecodeCloudeventData(c *v1.CloudEvent, protoMes proto.Message) error {
	if err := c.GetProtoData().UnmarshalTo(protoMes); err != nil {
		return err
	}

	return nil
	// if err := c.DataAs(protoMes); err == nil {
	// 	return nil
	// }

	// err := proto.Unmarshal(c.Data(), protoMes)
	// if err == nil {
	// 	return nil
	// }

	// bytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
	// if err != nil {
	// 	return err
	// }

	// if err := proto.Unmarshal(bytes, protoMes); err == nil {
	// 	return nil
	// }
}
