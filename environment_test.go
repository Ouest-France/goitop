package goitop

import (
	"os"
	"testing"
)

func TestClient_GetEnvironment(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.GetEnvironment(os.Getenv("GOITOP_ENV_NAME"))
	if err != nil {
		t.Fatal(err)
	}

	if id != os.Getenv("GOITOP_ENV_ID") {
		t.Fatalf("GetEnvironment() got %s; want %s", id, os.Getenv("GOITOP_ENV_ID"))
	}
}
