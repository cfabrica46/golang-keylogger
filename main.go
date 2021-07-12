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

type InputEvent struct {
	Time       syscall.Timeval
	Type, Code uint16
	Value      int32
}

func main() {
	log.SetFlags(log.Lshortfile)
	keyobardPath, err := findKeyboardDevice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyobardPath)

	f, err := os.OpenFile(keyobardPath, os.O_RDONLY, os.ModeCharDevice)
	if err != nil {
		log.Fatal(err)
	}
	//	syscall.Open(keyobardPath, syscall.O_RDONLY, syscall.AF_KEY)
	defer f.Close()

	buffer := make([]byte, unsafe.Sizeof(InputEvent{}))

	event := &InputEvent{}
	err = binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
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
	return "", nil
}
