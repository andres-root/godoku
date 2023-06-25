BIN_DIR = bin
BIN_NAME = godoku

ifeq ($(OS), Windows_NT)
	OS = windows
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	PACKAGE = $(shell Get-Content go.mod -head 1 | Foreach-Object { $$data = $$_ -split " "; "{0}" -f $$data[1]})
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
	BIN = ${BIN_NAME}.exe
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME),Darwin)
		OS = macos
	else ifeq ($(UNAME),Linux)
		OS = linux
	else
	$(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	BIN = ${BIN_NAME}
endif

build:
	go build -o ${BIN_DIR}/${BIN} .

bump: generate
	go get -u ./...

clean:
	${RM_RF_CMD} ${BIN_DIR}