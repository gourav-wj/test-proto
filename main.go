package main

import (
	"fmt"
	"log"
	"math"

	"github.com/golang/protobuf/proto"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

func main() {
	gourav := &Person{
		Name: "Gourav",
		Age:  23,
	}

	data, err := proto.Marshal(gourav)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	log.Print(string(data))

	newGourav := &Person{}
	err = proto.Unmarshal(data, newGourav)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Println(newGourav.Name)
	fmt.Println(newGourav.Age)
}
