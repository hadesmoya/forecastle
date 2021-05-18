<h1 align="center">Forecastle</h1>
<div align="center"><p>🌧🏰🌥</p></div>
<div align="center"><strong>Is a go package that uses</strong></div>
<div align="center"><code>OpenWeather API</code> to help you</div>
<div align="center">implement weather forecasts in your apps</div>
<div align="center">with the speed of light ⚡️</div>
<br>
<div align="center">
<a href="https://github.com/h4desune/forecastle">
<img src="https://img.shields.io/github/go-mod/go-version/h4desune/forecastle.svg" alt="go version">
</a>
</div>

## Table of contents
- [Installation](#installation)
- [Example](#example)
- [Documentation](#documentation) 
  - [getCurrentWeather()](#getcurrentweathercity-units-outputdetails-apikey)
- [License](#license)

## Installation
Use the [go get](https://golang.org/cmd/go/) command to install forecastle.

```bash
go get github.com/h4desune/forecastle
```

## Example

```go
import "github.com/h4desune/forecastle"

func main() {

forecastle.GetCurrentWeather(Moscow, metric, default, XXX)
// where XXX - is an actual APIKEY for openweathermap.org
// output is: It's 29.33° metric and a clear sky right now in the Moscow

}
```

## Documentation

Forecastle is a go package to access [OpenWeather API](https://openweathermap.org/api) fast and easy. You can call only a function and get the weather data with the specific options you need. Below you can see all the functions and their arguments.

### `getCurrentWeather(city, units, outputDetails, ApiKey)`

The function <code>getCurrentWeather()</code> takes a city name, the units of measurement, details of the output and your OpenWeather API Key.

<code>city</code> - open weather supports over 200,000 cities.

<code>units</code> - standard, metric, and imperial units are available.

<code>outputDetails</code> - specify how much data do you want, it can be default or extended.

<code>ApiKey</code> - your [OpenWeatherMaps](https://openweathermap.org/api) API Key.

As a result it shows a temperature and current state of the weather in the city.

TODO: examples

## License
[Apache-2.0 License](http://www.apache.org/licenses/LICENSE-2.0)
