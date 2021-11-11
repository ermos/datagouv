package datagouv

import (
	"context"
	"fmt"
	"net/http"
)

const (
	CommuneFormatJSON = "json"
	CommuneFormatGeoJSON = "geojson"
)

type Commune struct {
	Code            string   `json:"code"`
	Nom             string   `json:"nom"`
	CodesPostaux    []string `json:"codesPostaux"`
	CodeDepartement string   `json:"codeDepartement"`
	CodeRegion      string   `json:"codeRegion"`
	Departement     struct {
		Code       string `json:"code"`
		Nom        string `json:"nom"`
		CodeRegion string `json:"codeRegion"`
		Region     struct {
			Code string `json:"code"`
			Nom  string `json:"nom"`
		} `json:"region"`
	} `json:"departement"`
	Region struct {
		Code string `json:"code"`
		Nom  string `json:"nom"`
	} `json:"region"`
	Population int `json:"population"`
	Surface    int `json:"surface"`
	Centre     struct {
	} `json:"centre"`
	Contour struct {
	} `json:"contour"`
}

type CommuneParameters struct {
	// PostalCode allows to search by Postal Code
	PostalCode	string		`query:"codePostal"`
	// Lat allows to search by latitude (in degrees)
	Lat 	float64			`query:"lat"`
	// Lon allows to search by longitude (in degrees)
	Lon 	float64			`query:"lon"`
	// Name allows to search by commune name
	Name 	string			`query:"nom"`
	// Boost allows to set boost mode for Name field
	Boost 	string			`query:"boost"`
	// Code allows to search by commune code (insee)
	Code	string			`query:"code"`
	// DepartmentCode allows to search by department ode
	DepartmentCode string	`query:"codeDepartement"`
	// RegionCode allows to search by region code
	RegionCode	string		`query:"codeRegion"`
	// Fields allows to return specific list of fields
	// List: nom, code, codesPostaux, centre, surface, contour, codeDepartement, departement, codeRegion, region, population
	Fields []string			`query:"fields"`
	// Format is expected response format
	Format string			`query:"format"`
	// Geometry define geographic output
	Geometry string			`query:"geometry"`
	// Limit allows to limit number of result
	Limit 	int				`query:"limit"`
}

func GetCommune(params CommuneParameters) (c []Commune, err error) {
	return GetCommuneTx(context.Background(), params)
}

func GetCommuneTx(ctx context.Context, params CommuneParameters) (c []Commune, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", setCommuneURL(params), nil)
	if err != nil {
		return
	}

	err = runRequest(req, &c)
	return
}

func setCommuneURL(params CommuneParameters) string {
	base := fmt.Sprintf("%s/communes", geoApiURL)
	return buildUrl(base, parseParameters(params))
}