package datagouv

import (
	"testing"
)

func TestGetCommunesByDepartmentOK(t *testing.T) {
	c, err := GetCommunesByDepartment("72", CommunesByDepartmentParameters{
		Fields: []string{ "name", "code" , "codesPostaux" },
	})
	if err != nil {
		t.Error(err)
	}

	if len(c) == 0 {
		t.Error("no result ?")
		return
	}
}

func TestGetCommunesByDepartmentWithoutCode(t *testing.T) {
	_, err := GetCommunesByDepartment("", CommunesByDepartmentParameters{})
	if err == nil {
		t.Error("no error ?")
	}
}