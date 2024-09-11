package infrastructure

import (
	"SebStudy/pb"

	"google.golang.org/protobuf/proto"
)

func DecodeCloudeventData(c *pb.CloudEvent, protoMes proto.Message) error {
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
