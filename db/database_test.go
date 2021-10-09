package db_test

import (
	"appointy-task/db"
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	database := db.GetDatabase()
	defer db.Close(context.TODO())
	exptected := "instagram"
	actual := fmt.Sprint(database)
	if !strings.Contains(actual, exptected) {
		t.Errorf("Expected: %v, got %v", exptected, actual)
	}
}
