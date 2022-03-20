package main

import (
	"log"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

func toJSON(pb protoiface.MessageV1) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb protoiface.MessageV1) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}
