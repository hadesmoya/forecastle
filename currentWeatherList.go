package forecastle

type CurrentWeatherList struct {
    Message string               `json:"message"`
    Cod     int               `json:"cod"`
    Count   int              `json:"count"`
    List    []CurrentWeather `json:"list"`
}
