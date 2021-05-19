package forecastle

type SolarJson struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	List []struct {
		Radiation struct {
			Ghi   int `json:"ghi"`
			Dni   int `json:"dni"`
			Dhi   int `json:"dhi"`
			GhiCs int `json:"ghi_cs"`
			DniCs int `json:"dni_cs"`
			DhiCs int `json:"dhi_cs"`
		} `json:"radiation"`
		Dt int `json:"dt"`
	} `json:"list"`
}
