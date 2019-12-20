package goitop

import (
	"os"
	"reflect"
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

func TestClient_VMCRUD(t *testing.T) {

	c := NewClient(os.Getenv("GOITOP_ADDR"), os.Getenv("GOITOP_USER"), os.Getenv("GOITOP_PASSWORD"))
	id, err := c.CreateVM("goitop", "9", "5", "1318", "5", "yes", "1", "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = c.DeleteVM(id)
		if err != nil {
			t.Fatalf("Failed to delete VM with id %s", id)
		}
	}()

	if id == "" {
		t.Fatalf("Empty VM id returned after creation")
	}

	gotVM, err := c.GetVM(id)
	if err != nil {
		t.Fatal(err)
	}

	wantVM := VM{
		Name:                  "goitop",
		OrgID:                 "9",
		EnvID:                 "5",
		ClusterID:             "1318",
		ExploitationServiceID: "5",
		ID:                    id,
		Backup:                "yes",
		BackupID:              "1",
		Description:           "goitop_test",
	}

	eq := reflect.DeepEqual(gotVM, wantVM)
	if !eq {
		t.Fatalf("VM created and VM read are not equals: got %v  want %v", gotVM, wantVM)
	}

	err = c.UpdateVM(id, "goitop", "9", "4", "1318", "1", "no", "0", "")
	if err != nil {
		t.Fatal(err)
	}

	gotVM, err = c.GetVM(id)
	if err != nil {
		t.Fatal(err)
	}

	wantVM = VM{
		Name:                  "goitop",
		OrgID:                 "9",
		EnvID:                 "4",
		ClusterID:             "1318",
		ExploitationServiceID: "1",
		ID:                    id,
		Backup:                "no",
		BackupID:              "0",
		Description:           "goitop_test",
	}

	eq = reflect.DeepEqual(gotVM, wantVM)
	if !eq {
		t.Fatalf("VM created and VM read are not equals: got %v  want %v", gotVM, wantVM)
	}
}
