package ffrec

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

type FFRec struct {
	filename string
	instance *exec.Cmd
}

func New(filename string) FFRec {
	return FFRec{
		filename: filename,
	}
}

func (f *FFRec) Record() {

	c := []string{"-f", "pulse", "-i", "default", f.filename + ".mp3"}

	if runtime.GOOS == "darwin" {
		c = []string{"-f", "avfoundation", "-i", ":1", f.filename + ".mp3"}
	}

	f.instance = exec.Command("ffmpeg", c...)
	err := f.instance.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (f *FFRec) Stop() {

	f.instance.Process.Signal(os.Interrupt)
}
