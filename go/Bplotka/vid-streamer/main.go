package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/Bplotka/sgl"
	"github.com/blackjack/webcam"
)

const frameTimeout = 5 * time.Second

type FrameSizes []webcam.FrameSize

func (slice FrameSizes) Len() int {
	return len(slice)
}

//For sorting purposes
func (slice FrameSizes) Less(i, j int) bool {
	ls := slice[i].MaxWidth * slice[i].MaxHeight
	rs := slice[j].MaxWidth * slice[j].MaxHeight
	return ls < rs
}

//For sorting purposes
func (slice FrameSizes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Just for checking if the lib I chose, works.
func main() {
	log := sgl.New(os.Stdout)
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		log.WithErr(err).Fatal("Failed to open video.")
	}
	defer cam.Close()

	format_desc := cam.GetSupportedFormats()
	log.Info("Available formats: ")
	for f, val := range format_desc {
		log.Info("%d -> %s", f, val)

		frames := FrameSizes(cam.GetSupportedFrameSizes(f))
		sort.Sort(frames)

		for i, value := range frames {
			fmt.Fprintf(os.Stderr, "[%d] %s\n", i+1, value.GetString())
		}
	}
	_, _, _, err = cam.SetImageFormat(webcam.PixelFormat(1196444237), 640, 480)
	if err != nil {
		log.WithErr(err).Fatal("Failed to set image format.")
	}

	err = cam.StartStreaming()
	if err != nil {
		log.WithErr(err).Fatal("Failed to start streaming.")
	}
	// In seconds.
	err = cam.WaitForFrame(uint32(40))
	switch err.(type) {
	case nil:
	case *webcam.Timeout:
		log.WithErr(err).Error("Timeout on webcam")
	default:
		log.WithErr(err).Fatal("WaitForFrame err")
	}

	num := 0
	for {
		// In seconds.
		err = cam.WaitForFrame(uint32(frameTimeout.Seconds()))
		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			log.WithErr(err).Error("Timeout on webcam")
			continue
		default:
			log.WithErr(err).Fatal("WaitForFrame err")
		}

		frame, err := cam.ReadFrame()
		if err != nil {
			log.WithErr(err).Error("ReadFrame err")
			continue
		}

		log.Info("Frame len %d", len(frame))
		err = ioutil.WriteFile(fmt.Sprintf("out/%d.jpeg", num), frame, os.ModePerm)
		if err != nil {
			log.WithErr(err).Error("Failed to write file")
		}

		num += 1
		time.Sleep(100 * time.Millisecond)
	}
}
