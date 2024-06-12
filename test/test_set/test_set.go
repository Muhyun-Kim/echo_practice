package testset

import (
	"database/sql"
	"log"
	"my-echo-app/database"
	testdatabase "my-echo-app/test/test_database"
	"os"
	"path/filepath"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestMain(m *testing.M) {
	log.Println("Setting up test environment...")

	// .env 파일 로드
	// .env 파일 절대 경로 설정
	absPath, err := filepath.Abs("../../../.env")
	if err != nil {
		log.Fatalf("Error determining absolute path: %v", err)
	}
	log.Printf("Absolute path: %s", absPath)

	// .env 파일 로드
	err = godotenv.Load(absPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println(".env file loaded successfully")

	// Mock 데이터베이스 설정
	var db *sql.DB
	db, testdatabase.Mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	log.Println("sqlmock initialized successfully")

	// SELECT VERSION() 쿼리 설정
	testdatabase.Mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("mock_version"))

	dsn := "sqlmock_db_0"
	database.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:        dsn,
		DriverName: "mysql",
		Conn:       db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}
	log.Println("gorm database connection established successfully")

	// 데이터베이스 연결 확인
	if database.DB == nil {
		log.Fatalf("database.DB is still nil after initialization")
	}

	// 테스트 실행
	log.Println("Running tests...")
	code := m.Run()

	// Mock 데이터베이스 종료
	if err := testdatabase.Mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %s", err)
	}
	log.Println("All expectations were met")

	os.Exit(code)
}
