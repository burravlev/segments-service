package repositories_test

import (
	"testing"
)

func TestSegments_Create(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected opening a stub database connection", err)
	// }
	// defer db.Close()
	// dialector := postgres.New(postgres.Config{
	// 	DSN:        "sqlmock_db_0",
	// 	DriverName: "postgres",
	// 	Conn:       db,
	// })
	// gormDB, err := gorm.Open(dialector, &gorm.Config{})
	// if err != nil {
	// 	t.Errorf("Failed to open gorm db, got error: %s", err)
	// }
	// repository := repositories.SegmentRespository(gormDB)
	// segment := models.Segment{
	// 	ID:     1,
	// 	UserID: 12,
	// 	Name:   "test",
	// }
	// mock.MatchExpectationsInOrder(false)
	// mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"segments\"()"))
	// err = mock.ExpectationsWereMet()
	// if err != nil {
	// 	t.Errorf("Failed to meet expectations, got error: %v", err)
	// }
}
