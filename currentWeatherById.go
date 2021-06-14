package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeatherById(
	cityId int,
	appid string,
	units string,
	lang string,
) (*CurrentWeather, error) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%v&appid=%s&units=%s&lang=%s",
		cityId,
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

	var jsonHandler CurrentWeather

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		return nil, err
	}

	return &jsonHandler, nil

}
