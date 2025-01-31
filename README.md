# Inter-Process Communication Example

DON'T USE THIS IN PRODUCTION. This example does not contain any synchronization mechanisms. It is only intended to demonstrate the basic concept of IPC using shared memory.

![Demo](./demo.gif)

This repository demonstrates a simple implementation of inter-process communication (IPC) between Go and Node.js using shared memory through memory-mapped files.

The example consists of two programs:

* A Go server that writes data to shared memory
* A Node.js client that reads the data from shared memory

## Pre-requisites

* Go 1.16 or later
* Node.js 14 or later
* Unix-like operating system (Linux, macOS, etc.)

## Running the Example

1. Run the Go app:

```bash
go run main.go
```

2. Run the Node.js app:

```bash
node reader.js
```

## Our Paid Apps

* [tradeslyFX Forex AI Roboadvisor](https://play.google.com/store/apps/details?id=com.tradesly.tradeslyfx)

* [tradeslyPro Cryptocurrency AI Roboadvisor](https://play.google.com/store/apps/details?id=com.tradesly.tradeslypro)