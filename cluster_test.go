package goitop

import (
	"os"
	"testing"
)

func TestClient_GetCluster(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.GetCluster(os.Getenv("GOITOP_CLUSTER_NAME"))
	if err != nil {
		t.Fatal(err)
	}

	if id != os.Getenv("GOITOP_CLUSTER_ID") {
		t.Fatalf("GetCluster() got %s; want %s", id, os.Getenv("GOITOP_CLUSTER_ID"))
	}
}
