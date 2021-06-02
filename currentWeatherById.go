package forecastle

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func CurrentWeatherById(cityId int, unitsOfMeasurement string, outputType string, appid string, lang string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%v&appid=%s&units=%s&lang=%s",
		cityId,
		appid,
		unitsOfMeasurement,
		lang,
	)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: handle all (http) errors that may appear

	body, _ := ioutil.ReadAll(response.Body)

	var jsonHandler currentWeatherJson

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		log.Fatal(err)
	}

	var outputWeather = weatherConditionCodesMap[jsonHandler.Weather[0].Id]

	switch {
	case outputType == "onlyTemperature":
		fmt.Printf("%v°", jsonHandler.Main.Temp)
	case outputType == "temperatureAndDescription":
		fmt.Printf("%v° - %v", jsonHandler.Main.Temp, outputWeather)
	case outputType == "test":
		fmt.Printf("%v° - %v", jsonHandler.Main.Temp, jsonHandler.Weather[0].Description)
	}

}
