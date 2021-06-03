package forecastle

type CurrentWeatherJsonMux struct {
	Message string `json:"message"`
	Cod string `json:"cod"`
	Count int `json:"count"`
	List []CurrentWeatherJson `json:"list"`
}