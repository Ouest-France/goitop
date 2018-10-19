package goitop

import (
	"os"
	"testing"
)

func TestClient_GetBackup(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.GetBackup(os.Getenv("GOITOP_BACKUP_NAME"))
	if err != nil {
		t.Fatal(err)
	}

	if id != os.Getenv("GOITOP_BACKUP_ID") {
		t.Fatalf("GetBackup() got %s; want %s", id, os.Getenv("GOITOP_BACKUP_ID"))
	}
}
