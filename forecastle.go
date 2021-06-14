package forecastle

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// A struct of call arguments.

type Forecastle struct {
    city string
    cityID int
    appID string
    units string
    language string
}

// Helper functions to make methods contain less code.

func apiCall(url string) (*CurrentWeather, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    // TODO: handle all (http) errors that may appear.

    body, _ := ioutil.ReadAll(response.Body)

    var jsonHandler CurrentWeather

    err = json.Unmarshal(body, &jsonHandler)
    if err != nil {
        return nil, err
    }

    return &jsonHandler, nil
}

// The Beginning of Methods.

func (w *Forecastle) CurrentWeatherByCity() (*CurrentWeather, error) {
    var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
        w.city,
        w.appID,
        w.units,
        w.language,
    )

    return apiCall(url)
}

func(w *Forecastle) CurrentWeatherByID() (*CurrentWeather, error) {
    var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%v&appid=%s&units=%s&lang=%s",
        w.cityID,
        w.appID,
        w.units,
        w.language,
    )

    return apiCall(url)
}