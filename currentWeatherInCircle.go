package forecastle

// api.openweathermap.org/data/2.5/find?lat={lat}&lon={lon}&cnt={cnt}&appid={API key}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeatherInCircle(
	latitude float64,
	longitude float64,
	cnt int,
	appid string,
	units string,
	lang string,
) (*CurrentWeatherJsonMux, error) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/find?lat=%v&lon=%v&cnt=%v&appid=%s&units=%s&lang=%s",
		latitude,
		longitude,
		cnt,
		appid,
		units,
		lang,
	)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// TODO: handle all (http) errors that may appear

	body, _ := ioutil.ReadAll(response.Body)

	var jsonHandler CurrentWeatherJsonMux

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		return nil, err
	}

	return &jsonHandler, nil
}
