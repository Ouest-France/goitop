package goitop

import (
	"os"
	"testing"
)

func TestClient_GetAllVM(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	vms, err := c.GetAllVM()
	if err != nil {
		t.Fatal(err)
	}

	if len(vms) == 0 {
		t.Fatal("No VM returned")
	}
}

func TestClient_CreateVM(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.CreateVM("goitop", "9", "5", "1318", "5")
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatalf("Empty VM id returned after creation")
	}

	err = c.DeleteVM(id)
	if err != nil {
		t.Fatalf("Failed to delete VM with id %s", id)
	}
}
