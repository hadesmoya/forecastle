package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeatherWithinRectangle(
	longitudeLeft int,
	latitudeBottom int,
	longitudeRight int,
	latitudeTop int,
	zoom int,
	appid string,
	units string,
	lang string,
) (*CurrentWeatherJsonMux, error) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/box/city?bbox=%v,%v,%v,%v,%v&appid=%s&units=%s&lang=%s",
		longitudeLeft,
		latitudeBottom,
		longitudeRight,
		latitudeTop,
		zoom,
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
