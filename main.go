package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	keyobard, err := findKeyboardDevice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyobard)
}

func findKeyboardDevice() (string, error) {
	var deviceName string
	path := "/sys/class/input/event%d/device/name"

	for i := 0; i < 255; i++ {
		b, _ := ioutil.ReadFile(fmt.Sprintf(path, i))

		deviceName = strings.ToLower(string(b))
		if strings.Contains(deviceName, "keyboard") {
			return deviceName, nil
		}
	}
	if deviceName == "" {
		return "", errors.New("keyboard not found")
	}
	return "", nil
}
