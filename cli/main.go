package main

import (
	"math/rand"
	"time"

	pb "github.com/katallaxie/voskhod/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pb.VoskhodClientCommand.Execute()
}
