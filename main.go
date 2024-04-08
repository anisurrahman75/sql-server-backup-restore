package main

import (
	"context"
	_ "gocloud.dev/blob/azureblob"
	"log"
)

var ctx = context.Background()

func main() {
	blobfs()
}

func blobfs() {

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
