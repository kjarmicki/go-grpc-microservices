#!/bin/bash

export DATA_SOURCE_URL="root:pass@tcp(127.0.0.1:3300)/order"
export APPLICATION_PORT=3000
export ENV=development
go run cmd/main.go
