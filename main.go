package main

import (
	"time"

	"github.com/coral/ffrec/ffrec"
)

func main() {

	nr := ffrec.New("test")
	nr.Record()
	time.Sleep(10 * time.Second)
	nr.Stop()

}
