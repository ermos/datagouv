package datagouv

import (
	"testing"
)

func TestGetCommunesByDepartmentOK(t *testing.T) {
	c, err := GetCommunesByDepartment(CommunesByDepartmentParameters{
		Fields: []string{ "name", "code" , "codesPostaux" },
		Code: "72",
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
	_, err := GetCommunesByDepartment(CommunesByDepartmentParameters{})
	if err == nil {
		t.Error("no error ?")
	}
}