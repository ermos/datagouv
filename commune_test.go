package datagouv

import (
	"testing"
)

func TestGetCommune(t *testing.T) {
	c, err := GetCommune(CommuneParameters{
		Name: "La Flèche",
		Fields: []string{ "name", "code" , "codesPostaux" },
		Limit: 10,
	})
	if err != nil {
		t.Error(err)
	}

	if len(c) == 0 {
		t.Error("no result ?")
		return
	}

	if c[0].Nom != "La Flèche" {
		t.Errorf("expected name: La Flèche, got: %s", c[0].Nom)
	}
}