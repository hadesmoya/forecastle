package forecastle

type CurrentWeatherList struct {
    Message string               `json:"message"`
    Cod     string               `json:"cod"`
    Count   int              `json:"count"`
    List    []CurrentWeather `json:"list"`
}
