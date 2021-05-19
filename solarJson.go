package forecastle

type SolarJson struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	List []struct {
		Radiation struct {
			Ghi   float64 `json:"ghi"`
			Dni   float64 `json:"dni"`
			Dhi   float64 `json:"dhi"`
			GhiCs float64 `json:"ghi_cs"`
			DniCs float64 `json:"dni_cs"`
			DhiCs float64 `json:"dhi_cs"`
		} `json:"radiation"`
		Dt int `json:"dt"`
	} `json:"list"`
}
