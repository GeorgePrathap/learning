package serializertest_test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"gitlab.com/GeorgePrathap/protobuf/pb/pb"
	"gitlab.com/GeorgePrathap/protobuf/sample"
	"gitlab.com/GeorgePrathap/protobuf/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	if err != nil {
		fmt.Println("can't write the binary")
	}

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	if err != nil {
		fmt.Println("can't read the binary")
	}

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	if err != nil {
		fmt.Println("can't read the binary")
	}

	if proto.Equal(laptop1, laptop2) {
		fmt.Println("both match")
		return
	}

	fmt.Println("not match")
}
