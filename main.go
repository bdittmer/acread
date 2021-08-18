package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	cachePath := os.Getenv("ACREAD_PATH")
	files := os.Args[1:]

	if cachePath == "" {
		log.Fatal("ACREAD_PATH not set")
	}

	if len(files) == 0 {
		log.Fatal("no input files")
	}

	var jm *jsonpb.Marshaler
	for _, file := range files {
		filePath := filepath.Join(cachePath, file[:2], file)
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("cannot open %s: %s", filePath, err)
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("cannot read %s: %s", filePath, err)
		}

		var a repb.ActionResult
		if err := proto.Unmarshal(data, &a); err != nil {
			log.Fatalf("cannot unmarshal %s: %s", filePath, err)
		}

		if err := jm.Marshal(os.Stdout, &a); err != nil {
			log.Fatalf("cannot marshal %s to json: %s", filePath, err)
		}
	}
}
