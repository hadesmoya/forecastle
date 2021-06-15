package forecastle

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const baseURI = "https://api.openweathermap.org/data/2.5/"

type Client struct {
    appID string
}

func NewClient(appID string) *Client {
    return &Client{appID: appID}
}

// API call template function

func apiCall(url string) (body []byte, err error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(response.Body)
}

func unmarshalCurrentWeather(body []byte) (*CurrentWeather, error) {
    var jsonHandler CurrentWeather

    err := json.Unmarshal(body, &jsonHandler)
    if err != nil {
        return nil, err
    }

    return &jsonHandler, nil
}

// The Beginning of Methods Declaration.

func (client *Client) CurrentWeatherByCity(city, units, language string) (*CurrentWeather, error) {
    var url = fmt.Sprintf(baseURI+"weather?q=%s&appid=%s&units=%s&lang=%s",
        city,
        client.appID,
        units,
        language,
    )

    body, err := apiCall(url)
    if err != nil {
        return nil, err
    }

    return unmarshalCurrentWeather(body)
}

func (client *Client) CurrentWeatherByID(cityID int, units, language string) (*CurrentWeather, error) {
    var url = fmt.Sprintf(baseURI+"weather?id=%v&appid=%s&units=%s&lang=%s",
        cityID,
        client.appID,
        units,
        language,
    )

    body, err := apiCall(url)
    if err != nil {
        return nil, err
    }

    return unmarshalCurrentWeather(body)
}

func (client *Client) CurrentWeatherByGeo(latitude, longitude float64, units, language string) (*CurrentWeather, error) {
    var url = fmt.Sprintf(baseURI+"weather?lat=%v&lon=%v&appid=%s&units=%s&lang=%s",
        latitude,
        longitude,
        client.appID,
        units,
        language,
    )

    body, err := apiCall(url)
    if err != nil {
        return nil, err
    }

    return unmarshalCurrentWeather(body)
}

func (client *Client) CurrentWeatherByZip(zip int, countryCode string, units, language string) (*CurrentWeather, error) {
    var url = fmt.Sprintf(baseURI+"weather?zip=%v,%s&appid=%s&units=%s&lang=%s",
        zip,
        countryCode,
        client.appID,
        units,
        language,
    )

    body, err := apiCall(url)
    if err != nil {
        return nil, err
    }

    return unmarshalCurrentWeather(body)
}
