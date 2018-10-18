package goitop

import (
	"os"
	"testing"
)

func TestClient_GetOrganization(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.GetOrganization(os.Getenv("GOITOP_ORG_NAME"))
	if err != nil {
		t.Fatal(err)
	}

	if id != os.Getenv("GOITOP_ORG_ID") {
		t.Fatalf("GetOrganization() got %s; want %s", id, os.Getenv("GOITOP_ORG_ID"))
	}
}
