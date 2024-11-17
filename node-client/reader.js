const fs = require('fs');

const FILE_PATH = "/tmp/shm-ex01";
const SHM_SIZE = 1024;

// struct layout
/*
type SharedData struct {
    Message   [256]byte // 256 bytes
    Counter   int64     // 8 bytes
    Timestamp int64     // 8 bytes
}
*/
const MESSAGE_OFFSET = 0;
const MESSAGE_SIZE = 256;
const COUNTER_OFFSET = MESSAGE_SIZE;
const COUNTER_SIZE = 8;
const TIMESTAMP_OFFSET = COUNTER_OFFSET + COUNTER_SIZE;
const TIMESTAMP_SIZE = 8;

async function main() {
    while (!fs.existsSync(FILE_PATH)) {
        console.log("Waiting for the shared memory to be created...");
        await new Promise(resolve => setTimeout(resolve, 1000));
    }

    const fd = fs.openSync(FILE_PATH, 'r');
    const buffer = Buffer.alloc(SHM_SIZE);

    setInterval(() => {
        fs.readSync(fd, buffer, 0, SHM_SIZE, 0);
        const message = buffer.toString('utf8', MESSAGE_OFFSET, MESSAGE_OFFSET + MESSAGE_SIZE);
        const counter = buffer.readBigInt64LE(COUNTER_OFFSET);
        const timestamp = buffer.readBigInt64LE(TIMESTAMP_OFFSET);

        console.log(`Message: ${message}, Counter: ${counter}, Timestamp: ${timestamp}`);
    }, 2000);
}

main().catch(console.error);