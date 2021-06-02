package forecastle

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func CurrentWeatherById(cityId string, unitsOfMeasurement string, outputType string, appid string, lang string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s",
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
		fmt.Printf("%v°%s", jsonHandler.Main.Temp, unitsOfMeasurement)
	case outputType == "testMode":
		fmt.Printf("It's a %s right now in the %s. The temperature = %v° %s, but it feels like %v°", outputWeather, jsonHandler.Name,
			jsonHandler.Main.Temp, unitsOfMeasurement, jsonHandler.Main.FeelsLike)
	}

}
