package datagouv

import (
	"context"
	"net/http"
	"testing"
)

func TestBuildURL(t *testing.T) {
	s := buildUrl("http://localhost", map[string]string{
		"a": "test",
		"b": "10",
	})

	sExpected := "http://localhost?a=test&b=10"

	if s != sExpected {
		t.Errorf("url got : %s // url expected : %s", s, sExpected)
	}
}

func TestParseParameters(t *testing.T) {
	p := parseParameters(struct {
		WithQuery 		string		`query:"a"`
		WithoutQuery 	string
		WithoutValue 	string		`query:"b"`
		TypeInt 		int			`query:"int"`
		TypeFloat64 	float64		`query:"float64"`
		TypeSliceString	[]string	`query:"slice_string"`
	}{
		WithQuery: "france",
		TypeInt: 10,
		TypeFloat64: 10.4,
		TypeSliceString: []string{ "a" },
	})

	if p["a"] != "france" {
		t.Error("a is not equal to france")
	}
}

func TestRunRequestOK(t *testing.T) {
	var res []interface{}
	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"https://geo.api.gouv.fr/communes?limit=1",
		nil,
		)
	if err != nil {
		t.Error(err)
		return
	}

	err = runRequest(req, &res)
	if err != nil {
		t.Error(err)
	}
}

func TestRunRequestResponseNot200(t *testing.T) {
	var res []interface{}
	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"https://geo.api.gouv.fr/api-not-exist",
		nil,
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = runRequest(req, &res)
	if err == nil {
		t.Error("err is nil ?")
	}
}

func TestRunRequestClientFailed(t *testing.T) {
	var res []interface{}
	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		"https://url-not-exist",
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	err = runRequest(req, &res)
	if err == nil {
		t.Error("err is nil ?")
	}
}

//func TestFetchOk(t *testing.T) {
//	var api []API
//	err := Fetch("test/fetch", &api, _toApiFake)
//	if err != nil {
//		t.Errorf(err.Error())
//	}
//	for _, a := range api {
//		if len(a.Routes) != 1 {
//			t.Errorf("cannot get %s's route", a.Controller)
//		}
//	}
//}