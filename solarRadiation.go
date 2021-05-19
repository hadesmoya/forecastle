package forecastle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// http://api.openweathermap.org/data/2.5/solar_radiation?lat={lat}&lon={lon}&appid={API key}

func SolarRadiation(lat, lon float64, appid string) {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/solar_radiation?lat=%v&lon=%v&appid=%s",
		lat,
		lon,
		appid,
	)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: handle all (http) errors that may appear

	body, _ := ioutil.ReadAll(response.Body)

	var SolarHandler SolarJson

	err = json.Unmarshal(body, &SolarHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latitude = %v; Longtitude = %v\n. GHI = %v",
		SolarHandler.Coord.Lat,
		SolarHandler.Coord.Lon,
		SolarHandler.List[0].Radiation.Ghi)
}