package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetCurrentWeather function shows you the current weather in more than 200.000 cities all over the world

func GetCurrentWeather(city string, unitsOfMeasurement string, outputType string, appid string, lang string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
		city,
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

	var output = jsonHandler.Weather[0].Description

	switch {
	case outputType == "ShowTemperature":
		fmt.Printf("%v°", jsonHandler.Main.Temp)
	case outputType == "ShowTemperatureAndDescription":
		fmt.Printf("%v° - %v", jsonHandler.Main.Temp, output)
	}

}
