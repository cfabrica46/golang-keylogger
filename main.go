package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	keyobardPath, err := findKeyboardDevice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keyobardPath)
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
