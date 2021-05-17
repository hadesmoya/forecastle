package forecastle

type jsonData struct {
	Coord      CoordJson     `json:"coord"`
	Weather    []WeatherJson `json:"weather"`
	Main       MainJson      `json:"main"`
	Wind       WindJson      `json:"wind"`
	Clouds     CloudsJson    `json:"clouds"`
	Sys        SysJson       `json:"sys"`
	Visibility int           `json:"visibility"`
	Base       string        `json:"base"`
	Dt         int           `json:"dt"`
	Timezone   int           `json:"timezone"`
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Cod        int           `json:"cod"`
}

type CoordJson struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherJson struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainJson struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type CloudsJson struct {
	All int `json:"all"`
}

type SysJson struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type WindJson struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  int     `json:"gust"`
}

