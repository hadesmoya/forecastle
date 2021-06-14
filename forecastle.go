package forecastle

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Forecastle struct {
    city string
    appID string
    units string
    language string
}

func apiCall(url string) (*CurrentWeather, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    // TODO: handle all (http) errors that may appear

    body, _ := ioutil.ReadAll(response.Body)

    var jsonHandler CurrentWeather

    err = json.Unmarshal(body, &jsonHandler)
    if err != nil {
        return nil, err
    }

    return &jsonHandler, nil
}

func (w *Forecastle) currentWeatherInCity() (*CurrentWeather, error) {
    var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
        w.city,
        w.appID,
        w.units,
        w.language,
    )

    return apiCall(url)
}

