package main

import (
	"fmt"
	"io/ioutil"
	"os"

	repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/golang/protobuf/proto"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var a repb.ActionResult
	if err := proto.Unmarshal(data, &a); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", proto.MarshalTextString(&a))
}
