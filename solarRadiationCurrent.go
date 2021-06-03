package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SolarRadiationCurrent(latitude, longitude float64, appid string) (*SolarJson, error) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/solar_radiation?lat=%v&lon=%v&appid=%s",
		latitude,
		longitude,
		appid,
	)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// TODO: handle all (http) errors that may appear

	body, _ := ioutil.ReadAll(response.Body)

	var solarHandler SolarJson

	err = json.Unmarshal(body, &solarHandler)
	if err != nil {
		return nil, err
	}

	return &solarHandler, nil
}

