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
<a href="https://github.com/h4desune/forecastle">
<img src="https://img.shields.io/github/stars/h4desune/forecastle?style=social" alt="github stars"> 
</a>
</div>

## Table of contents
- [Installation](#installation)
- [Updating](#updating)
- [Example](#example)
- [Documentation](#documentation) 
  - [GetCurrentWeather()](#getcurrentweathercity-units-outputdetails-apikey)
  - [SolarRadiationCurrent()](#solarradiationcurrent)
- [License](#license)

## Installation
Use the [go get](https://golang.org/cmd/go/) command to install the latest forecastle release.
```bash
go get github.com/h4desune/forecastle
```
See the [Updating](#updating) section to learn how to update or downgrade the package.

## Updating

Update package to the last version by using <code>-u</code> flag:
```bash
go get -u github.com/h4desune/forecastle
```

Update package to the specific release version by using <code>-u</code> flag and release version after <code>@</code>:

```bash
go get -u github.com/h4desune/forecastle@vX.X.X
```
There are two examples for <code>v0.0.11</code> and <code>2.12.3</code>:
```bash
go get -u github.com/h4desune/forecastle@v.0.0.11
```
```bash
go get -u github.com/h4desune/forecastle@v.2.12.3
```

Check [the releases page](https://github.com/h4desune/forecastle/releases) if you can not decide which one will fit all your needs. I always recommend using the latest one.

## Features

Right now, forecastle can help you in making calls to the next OpenWeather APIs:
- [Current weather data](https://openweathermap.org/current)
- [Solar Radiation](https://openweathermap.org/api/solar-radiation)

I'm planning to make this list bigger in the future.

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

### `GetCurrentWeather(city, units, outputDetails, ApiKey)`

The function <code>GetCurrentWeather()</code> takes a city name, the units of measurement, details of the output and your OpenWeather API Key.

<code>city</code> - open weather supports over 200,000 cities.

<code>units</code> - standard, metric, and imperial units are available.

<code>outputDetails</code> - specify how much data do you want, it can be default or extended.

<code>appid</code> - is an API Key for OpenWeather

As a result it shows a temperature and current state of the weather in the city.

TODO: examples


***

### `SolarRadiationCurrent(lat, lon, appid)`

The function <code>SolarRadiationCurrent()</code> takes a latitude, longitude and your OpenWeather API Key.

<code>lat</code> - (latitude) is a float64 number, for example: 32.724319

<code>lon</code> - (longitude) is a float64 number. for example: -114.624397

<code>appid</code> - is an API Key for OpenWeather

It provides you with current [solar radiation](https://www.energy.gov/eere/solar/solar-radiation-basics) data for any coordinates on the globe.

[See OpenWeatherMap DOCS for more](https://openweathermap.org/api/solar-radiation#concept)



## License
[Apache-2.0 License](http://www.apache.org/licenses/LICENSE-2.0)
