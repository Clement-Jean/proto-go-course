package main

import (
	"fmt"
	"reflect"

	pb "github.com/Clement-Jean/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
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
		fmt.Printf("This is an Id: %d\n", message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Printf("This is a message: %s\n", message.(*pb.Result_Message).Message)
	default:
		fmt.Printf("message has unexpected type: %T\n", x)
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
	return message
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
		//EyeColor: 1,
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	sm2 := &pb.Simple{}
	readFromFile(path, sm2)
	fmt.Println("Read the content:", sm2)
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// doOneOf(&pb.Result_Id{Id: 42})
	// doOneOf(&pb.Result_Message{Message: "My name"})
	// fmt.Println(doMap())
	// doFile(doSimple())
	// fmt.Println(doFromJSON(toJSON(doSimple()), reflect.TypeOf(pb.Simple{})))
	// fmt.Println(doFromJSON(toJSON(doComplex()), reflect.TypeOf(pb.Complex{})))
	// fmt.Println(doFromJSON(`{"id": 42, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}
