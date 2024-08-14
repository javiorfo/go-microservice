package dummy_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/repository"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var container testcontainers.Container
var repo repository.DummyRepository

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	var err error
	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start container: %s", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %s", err)
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get container port: %s", err)
	}

	dsn := "host=" + host + " port=" + port.Port() + " user=testuser password=testpass dbname=testdb sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err := db.AutoMigrate(&model.Dummy{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	repo = repository.NewDummyRepository(db)

	// Run the tests
	code := m.Run()

	// Cleanup
	if err := container.Terminate(ctx); err != nil {
		log.Fatalf("Failed to terminate container: %s", err)
	}

	os.Exit(code)
}

func TestDummy(t *testing.T) {

	dummyRecord := model.Dummy{Info: "testname"}

	if err := repo.Create(&dummyRecord); err != nil {
		t.Fatalf("Failed to insert record: %v", err)
	}

	dummy, err := repo.FindById(dummyRecord.ID)
	if err != nil {
		t.Fatalf("Failed to query record: %v", err)
	}

	if dummy.Info != "testname" {
		t.Errorf("Expected name to be 'testname', got '%s'", dummy.Info)
	}
}
