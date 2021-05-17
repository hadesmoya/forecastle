package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
getCurrentWeather function shows you the current weather in more than 200.000 cities all over the world
*/
func GetCurrentWeather(city string, unitsOfMeasurement string, outputDetails string, ApiKey string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s",
		city,
		ApiKey,
		unitsOfMeasurement,
	)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: handle all (http) errors that may appear

	// fmt.Println(url)
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(response)

	var jsonHandler jsonData

	err = json.Unmarshal(body, &jsonHandler)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v", jsonHandler)

	var outputWeather = weatherConditionCodesMap[jsonHandler.Weather[0].Id]  // uses the value of weatherConditionCodesMap as a current weather state output

	switch {

	case outputDetails == "default":
		fmt.Printf("It's %v° %s and a %s right now in the %s", jsonHandler.Main.Temp, unitsOfMeasurement, outputWeather, jsonHandler.Name)
	case outputDetails == "extended":
		fmt.Printf("It's a %s right now in the %s. The temperature = %v° %s, but it feels like %v°", outputWeather, jsonHandler.Name,
			jsonHandler.Main.Temp, unitsOfMeasurement, jsonHandler.Main.FeelsLike)
	}

}
