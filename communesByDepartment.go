package datagouv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type CommunesByDepartmentParameters struct {
	// Fields allows to return specific list of fields
	// List: nom, code, codesPostaux, centre, surface, contour, codeDepartement, departement, codeRegion, region, population
	Fields []string			`query:"fields"`
	// Format is expected response format
	Format string			`query:"format"`
	// Geometry define geographic output
	Geometry string			`query:"geometry"`
}

func GetCommunesByDepartment(code string, params CommunesByDepartmentParameters) (c []Commune, err error) {
	return GetCommunesByDepartmentTx(context.Background(), code, params)
}

func GetCommunesByDepartmentTx(ctx context.Context, code string, params CommunesByDepartmentParameters) (c []Commune, err error) {
	if code == "" {
		err = errors.New("code parameter is required")
		return
	}

	req, err := http.NewRequestWithContext(ctx, "GET", setCommunesByDepartmentURL(code, params), nil)
	if err != nil {
		return
	}

	err = runRequest(req, &c)
	return
}

func setCommunesByDepartmentURL(code string, params CommunesByDepartmentParameters) string {
	base := fmt.Sprintf("%s/departements/%s/communes", geoApiURL, code)
	return buildUrl(base, parseParameters(params))
}