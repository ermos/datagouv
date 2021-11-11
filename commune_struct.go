package datagouv

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
