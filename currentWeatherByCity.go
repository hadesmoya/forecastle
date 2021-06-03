package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeatherByCity(
	city string,
	appid string,
	units string,
	lang string,
) (*CurrentWeatherJson, error) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
		city,
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

	var jsonHandler CurrentWeatherJson

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		return nil, err
	}

	return &jsonHandler, nil
}
