package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"

	complexpb "github.com/simplesteph/protobuf-example-go/src/complex"
	enumpb "github.com/simplesteph/protobuf-example-go/src/enum_example"
	simplepb "github.com/simplesteph/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	readandWriteDemo(sm)
	//jsonDemo(sm)
	//doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	eb := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}
	fmt.Println(eb)
}

func jsonDemo(pb proto.Message) {
	out := toJSON(pb)
	fmt.Println("JSON", out)
	sm := &simplepb.SimpleMessage{}
	fromJSON(out, sm)
	fmt.Println("Reconstructed from json", sm)
}

func toJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	jsonpb.UnmarshalString(in, pb)
}

func readandWriteDemo(pb proto.Message) {
	writeToFile("simple.bin", pb)
	sm := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm)
	fmt.Println(sm)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cant write to file", err)
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong reading the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Something went wrong while reconstructing struct", err)
		return err
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple",
		SampleList: []int32{1, 4, 7, 8},
	}
	sm.Name = "I renamed you"
	fmt.Println(sm)
	return &sm
}
