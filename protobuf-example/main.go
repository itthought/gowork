package main

import (
	"fmt"
	"github.com/itthought/gowork/protobuf-example/simple"

	"github.com/golang/protobuf/proto"
	"log"
	"io/ioutil"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	smt := doSimple()
	fmt.Println(smt)
	readWriteProtoDemo(smt)
	jsonDemo(smt)

}

func jsonDemo(message proto.Message) {
	json, err := toJSON(message)
	if err==nil {
		fmt.Println("Json string is ",json)
		sm2 := &simplepb.SimpleMessage{}
		fromJSON(json, sm2)

		fmt.Println("Proto struct ",sm2)
	}
}

func readWriteProtoDemo(message proto.Message){
	writeToFile("simple.bin",message)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin",sm2)
	fmt.Println("Reading from file :",sm2)

}

func toJSON(message proto.Message)  (string, error) {

	marshlar := jsonpb.Marshaler{}
	out , err :=marshlar.MarshalToString(message)
	if err!=nil {
		log.Fatal("Can't convert to json :",err )
		return out, err
	}

	return  out, nil
}

func fromJSON(json string, message proto.Message)  {

	err :=jsonpb.UnmarshalString(json, message)
	if err!=nil {
		log.Fatal("Can't convert json to object :",err )
	}

}

func writeToFile(fn string, message proto.Message) error {

	out, err := proto.Marshal(message)
	if nil!=err {
		log.Fatal("Can't serialize object :",err )
		return err
	}

	if err := ioutil.WriteFile(fn, out, 0644); err !=nil {
		log.Fatal("Can't write data to file :",err )
		return err
	}

	fmt.Println("Data has been written to file")
	return nil

}

func readFromFile(fn string, pb proto.Message) error{

	io , err :=ioutil.ReadFile(fn)
	if err!=nil {
		log.Fatal("Can't read data from file :",err )
		return err
	}

	err2 := proto.Unmarshal(io,pb)
	if err2!=nil{
		log.Fatal("Can't deserialized object :",err )
		return err
	}

	return nil

}
func doSimple() *simplepb.SimpleMessage {
	smt := simplepb.SimpleMessage{
		Name:       "radhey",
		Id:         1211,
		IsSimple:   false,
		SampleList: []int32{1, 2, 3, 4, 5},
	}
	fmt.Println(smt)
	return &smt
}
