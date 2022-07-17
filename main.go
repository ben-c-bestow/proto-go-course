package main

import (
	"fmt"
	"reflect"

	pb "github.com/ben-c-bestow/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "my name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "other name"},
			{Id: 44, Name: "Bob"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("Unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"arthur": {Id: 42},
			"ted":    {Id: 69},
			"nick":   {Id: 666},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	// fmt.Println(jsonString)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func doAddressBook() *pb.AddressBook {
	return &pb.AddressBook{
		People: []*pb.Person{
			{
				Name:  "Arthur Dent",
				Id:    42,
				Email: "dontpanic@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "+44 1234 567890", Type: pb.Person_MOBILE},
					{Number: "+44 4444 444444", Type: pb.Person_WORK},
				},
			},
		},
	}
}

func main() {
	// doOneOf(&pb.Result_Id{Id: 42})
	// doOneOf(&pb.Result_Message{Message: "hi"})
	// fmt.Println(doMap())

	// doFile(doSimple())
	// jsonString := doToJSON(doComplex())
	// message := doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	// fmt.Println(jsonString)
	// fmt.Println(message)

	jsonString := doToJSON(doAddressBook())
	fmt.Println(jsonString)
}
