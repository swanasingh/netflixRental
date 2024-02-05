package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
	"time"
)

func TestContainer(m *testing.T) {
	ctx := context.Background()

	// Define a PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5434/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "pass",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithStartupTimeout(time.Minute * 2),
	}

	// Create the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Println("could not connect to container")
		log.Fatal(err)
	}
	defer container.Terminate(ctx)

	// Get container host
	host, err := container.Host(ctx)
	if err != nil {
		fmt.Println("could not get container host")
		log.Fatal(err)
	}

	// Get container port
	port, err := container.MappedPort(ctx, "5434")
	if err != nil {
		fmt.Println("could not get container port")
		log.Fatal(err)
	}
	fmt.Println("host ", host)
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=root password=pass dbname=netflix-rental sslmode=disable",
		host, port.Port())

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("could not connect to PostgreSQL")
		log.Fatal(err)
	}
	defer db.Close()

	// Your database test code here
	// Run tests
	//code := m.Run()

	// Clean up
	if err := container.Terminate(ctx); err != nil {
		log.Fatal(err)
	}

	// Exit
	//os.Exit(code)
}
