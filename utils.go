package datagouv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	geoApiURL = "https://geo.api.gouv.fr"
)

type errorMessage struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func buildUrl(base string, query map[string]string) string {
	v := url.Values{}

	for key, value := range query {
		v.Set(key, value)
	}

	return base + "?" + v.Encode()
}

func parseParameters(p interface{}) (res map[string]string) {
	l := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	res = make(map[string]string)

	for i := 0; i < l.NumField(); i++ {
		var value string

		field := l.Field(i)

		if field.Tag.Get("query") == "" {
			continue
		}

		switch field.Type.String() {
		case "string":
			value = v.Field(i).Interface().(string)
			break
		case "int":
			valueInterface := v.Field(i).Interface().(int)
			if valueInterface != 0 {
				value = fmt.Sprintf("%d", valueInterface)
			}
			break
		case "float64":
			valueInterface := v.Field(i).Interface().(float64)
			if valueInterface != 0 {
				value = fmt.Sprintf("%f", valueInterface)
			}
			break
		case "[]string":
			valueInterface := v.Field(i).Interface().([]string)
			if len(valueInterface) != 0 {
				value = strings.Join(valueInterface, ",")
			}
			break
		}

		if value != "" {
			res[field.Tag.Get("query")] = value
		}
	}

	return
}

func runRequest(req *http.Request, v interface{}) (err error) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		var errStruct errorMessage

		err = json.Unmarshal(body, &errStruct)
		if err != nil {
			return fmt.Errorf("api error (%d): not found", res.StatusCode)
		}

		return fmt.Errorf("api error (%d): %s", errStruct.Code, errStruct.Message)
	}

	return json.Unmarshal(body, &v)
}