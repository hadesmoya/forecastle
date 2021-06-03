package forecastle

// api.openweathermap.org/data/2.5/find?lat={lat}&lon={lon}&cnt={cnt}&appid={API key}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeatherInCircle(
	lat float64,
	lon float64,
	cnt int,
	appid string,
	units string,
	lang string,
) (*CurrentWeatherJson, error) {
	var url = fmt.Sprintf("api.openweathermap.org/data/2.5/find?lat=%v&lon=%v&cnt=%v&appid=%s&units=%s&lang=%s",
		lat,
		lon,
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

	var jsonHandler CurrentWeatherJson

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		return nil, err
	}

	return &jsonHandler, nil
}
