PACKAGE_NAME = cmd

BINARY_NAME = cmd

ENTRY_POINT = cmd/main.go

GO = go

# Set the build flags
BUILD_FLAGS = -v

# Set the run command
RUN_COMMAND = $(GO) run $(ENTRY_POINT)

# Default target: build the project
build:
	go build ./cmd/main.go

# Target: run the project
run:
	$(RUN_COMMAND)

# Target: clean the project
clean:
	rm -f $(BINARY_NAME)

# Target: build and run the project
all: build run

.PHONY: build run clean all