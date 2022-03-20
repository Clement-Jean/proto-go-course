package main

import (
	"fmt"

	pb "github.com/Clement-Jean/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		Name:        "My name",
		IsSimple:    true,
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "My name 2"},
			{Id: 44, Name: "My name 3"},
		},
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id))
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message))
	default:
		fmt.Errorf("message has unexpected type %v", x)
	}
}

func doMap() *pb.MapExample {
	message := &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}

	fmt.Println(message)
	return message
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
		//EyeColor: 1,
	}
}

func doJson(p protoiface.MessageV1) {
	jsonString := toJSON(p)
	fmt.Println(jsonString)

	sm2 := &pb.Simple{}
	fromJSON(jsonString, sm2)
	fmt.Println(sm2)
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	sm2 := &pb.Simple{}
	readFromFile(path, sm2)
	fmt.Println("Read the content:", sm2)
}

func main() {
	fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// doOneOf(&pb.Result_Id{Id: 42})
	// doOneOf(&pb.Result_Message{Message: "My name"})
	// doMap()
	// doJson(doSimple())
	// doFile(doSimple())
}
