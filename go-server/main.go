package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	FILE_PATH = "/tmp/shm-ex01"
	SHM_SIZE  = 1024
)

type SharedData struct {
	Message   [256]byte // 256 bytes
	Counter   int64     // 8 bytes
	Timestamp int64     // 8 bytes
}

func main() {
	shutdown := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fd, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	defer os.Remove(FILE_PATH)

	if err := os.Truncate(FILE_PATH, SHM_SIZE); err != nil {
		panic(err)
	}

	data, err := unix.Mmap(int(fd.Fd()), 0, SHM_SIZE, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	defer unix.Munmap(data)

	sharedData := (*SharedData)(unsafe.Pointer(&data[0]))

	go func() {
		<-sigChan
		fmt.Println("Exiting...")
		close(shutdown)
	}()

	counter := int64(0)
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-shutdown:
			return
		case <-ticker.C:
			counter++
			sharedData.Counter = counter
			sharedData.Timestamp = time.Now().UnixNano()
			copy(sharedData.Message[:], fmt.Sprintf("Hello, IPC! %d", counter))

			fmt.Printf("Counter: %d, Timestamp: %d, Message: %s\n", sharedData.Counter, sharedData.Timestamp, sharedData.Message)

			time.Sleep(2 * time.Second)
		}
	}

}
