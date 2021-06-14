package forecastle

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Client struct {
    appID string
}

func NewClient(appID string) *Client {
    return &Client{appID: appID}
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

// TODO: const baseURI

// The Beginning of Methods Declaration.

func (client *Client) CurrentWeatherByCity(city string, units, language string) (*CurrentWeather, error) {
    var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
        city,
        client.appID,
        units,
        language,
    )

    return apiCall(url)
}

/*func(w *Forecastle) CurrentWeatherByID() (*CurrentWeather, error) {
    var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%v&appid=%s&units=%s&lang=%s",
        w.CityID,
        w.AppID,
        w.Units,
        w.Language,
    )

    return apiCall(url)
}*/