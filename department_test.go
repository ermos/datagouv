package datagouv

import (
	"testing"
)

func TestGetDepartmentOK(t *testing.T) {
	c, err := GetDepartment(DepartmentParameters{
		Code: "72",
	})
	if err != nil {
		t.Error(err)
	}

	if len(c) == 0 {
		t.Error("no result ?")
		return
	}

	if c[0].Nom != "Sarthe" {
		t.Errorf("expected name: Sarthe, got: %s", c[0].Nom)
	}
}