package datagouv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type CommunesByDepartmentParameters struct {
	// Code allows to search by department code (Required)
	Code	string			`query:"code"`
	// Fields allows to return specific list of fields
	// List: nom, code, codesPostaux, centre, surface, contour, codeDepartement, departement, codeRegion, region, population
	Fields []string			`query:"fields"`
	// Format is expected response format
	Format string			`query:"format"`
	// Geometry define geographic output
	Geometry string			`query:"geometry"`
}

func GetCommunesByDepartment(params CommunesByDepartmentParameters) (c []Commune, err error) {
	return GetCommunesByDepartmentTx(context.Background(), params)
}

func GetCommunesByDepartmentTx(ctx context.Context, params CommunesByDepartmentParameters) (c []Commune, err error) {
	if params.Code == "" {
		err = errors.New("code parameter is required")
		return
	}

	req, err := http.NewRequestWithContext(ctx, "GET", setCommunesByDepartmentURL(params), nil)
	if err != nil {
		return
	}

	err = runRequest(req, &c)
	return
}

func setCommunesByDepartmentURL(params CommunesByDepartmentParameters) string {
	base := fmt.Sprintf("%s/departements/%s/communes", geoApiURL, params.Code)
	params.Code = ""
	return buildUrl(base, parseParameters(params))
}