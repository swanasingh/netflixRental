PACKAGE_NAME = cmd

BINARY_NAME = myapp

ENTRY_POINT = cmd/main.go

GO = go

# Set the build flags
BUILD_FLAGS = -o

# Set the run command
RUN_COMMAND = $(GO) run $(ENTRY_POINT)

# folder for binary
BINARY_FOLDER = out
# Default target: build the project
build:
	$(GO) build $(BUILD_FLAGS) $(BINARY_FOLDER)/$(BINARY_NAME) ./$(ENTRY_POINT)


# Target: run the project
run:
	$(RUN_COMMAND)

# Target: clean the project
clean:
	rm -f $(BINARY_FOLDER)/$(BINARY_NAME)

migrate:
	 migrate -path database/migration/ -database "postgresql://root:pass@localhost:5432/netflix-rental?sslmode=disable" -verbose up

# Target: build and run the project
all: build run

.PHONY: build run clean all