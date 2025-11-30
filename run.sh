#!/bin/bash

set -x

export DIVELOG_DBFILE_PATH=/home/andrija/backups/subsurface-2025-12.xml
export DIVELOG_MODE=dev

go run ./main.go &
SERVER_PID=$!

# Forward CTRL-C to the server process
trap "kill -INT $SERVER_PID" INT

sleep 0.5

firefox -new-tab http://127.0.0.1:8072/hms/dives &

wait $SERVER_PID
