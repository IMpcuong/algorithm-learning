package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func isValidFashion(size int, buttons []int) bool {
	if size != len(buttons) {
		return false
	}
	if size == 1 {
		return buttons[0] == 1
	}
	fCounter := 0
	for _, b := range buttons {
		if b == 1 {
			fCounter++
		}
	}
	return fCounter == size-1
}

const (
	VALID   string = "YES"
	INVALID string = "NO"
)

func main() {
	// https: //zetcode.com/golang/readinput/
	reader := bufio.NewReader(os.Stdin)

	// Create a channel to receive interrupt signal.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to indicate when to stop reading input.
	stopReading := make(chan bool)

	// Goroutine to listen for interrupt signal.
	go func() {
		<-interrupt
		fmt.Println("\nCtrl+C received. Exiting...")
		stopReading <- true
	}()

	stdinData := []string{}
	for i := 0; i < 2; i++ {
		select {
		case <-stopReading:
			return
		default:
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			stdinData = append(stdinData, strings.TrimRight(line, "\n"))
		}
	}

	size, _ := strconv.Atoi(stdinData[0])
	buttons := func(data []string) []int {
		buttons := []int{}
		for _, bStr := range data {
			bInt, _ := strconv.Atoi(bStr)
			buttons = append(buttons, bInt)
		}
		return buttons
	}(strings.Split(stdinData[1], " "))

	if isValidFashion(size, buttons) {
		fmt.Println(VALID)
	} else {
		fmt.Println(INVALID)
	}
}
