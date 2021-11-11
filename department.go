package datagouv

import (
	"context"
	"fmt"
	"net/http"
)

type Department struct {
	Code       string `json:"code"`
	Nom        string `json:"nom"`
	CodeRegion string `json:"codeRegion"`
	Region     struct {
		Code string `json:"code"`
		Nom  string `json:"nom"`
	} `json:"region"`
}

type DepartmentParameters struct {
	// Name allows to search by department name
	Name 	string			`query:"nom"`
	// Code allows to search by department code
	Code	string			`query:"code"`
	// RegionCode allows to search by region code
	RegionCode	string		`query:"codeRegion"`
	// Fields allows to return specific list of fields
	// List: nom, code, codeRegion, region
	Fields []string			`query:"fields"`
	// Limit allows to limit number of result
	Limit 	int				`query:"limit"`
}

func GetDepartment(params DepartmentParameters) (c []Department, err error) {
	return GetDepartmentTx(context.Background(), params)
}

func GetDepartmentTx(ctx context.Context, params DepartmentParameters) (c []Department, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", setDepartmentURL(params), nil)
	if err != nil {
		return
	}

	err = runRequest(req, &c)
	return
}

func setDepartmentURL(params DepartmentParameters) string {
	base := fmt.Sprintf("%s/departements", geoApiURL)
	return buildUrl(base, parseParameters(params))
}