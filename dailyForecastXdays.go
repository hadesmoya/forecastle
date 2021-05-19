package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DailyForecastXdays(city string, numberOfDays int, unitsOfMeasurement string, outputDetails string, ApiKey string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast/daily?q=%s&cnt=%v&appid=%s&units=%s",
		city,
		numberOfDays,
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

	var outputWeather = weatherConditionCodesMap[jsonHandler.Weather[0].Id]  // uses the value of weatherConditionCodesMap as a current weather state output

	switch {

	case outputDetails == "default":
		fmt.Printf(outputWeather, numberOfDays)
	case outputDetails == "extended":
		fmt.Printf("a")

	}

}
