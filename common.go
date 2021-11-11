package datagouv

const (
	FormatJSON = "json"
	FormatGeoJSON = "geojson"
)

type common struct {
	Name 	string			`query:"nom"`
	Code	string			`query:"code"`
	Fields []string			`query:"fields"`
}