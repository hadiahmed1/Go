package main

import (
	"fmt"

	pb "github.com/hadi/myapp/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   false,
		Name:       "Hadi",
		SampleList: []int32{2, 3, 4},
	}
}

func doComplex() *pb.Class {
	return &pb.Class{
		Teacher: &pb.Person{
			Id:   1,
			Name: "Some Teacher name",
		},
		Students: []*pb.Person{
			{Id: 2, Name: "Student One"},
			{Id: 2, Name: "Student Two"},
		},
	}
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
}
