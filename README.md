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

## Example

`Current weather in Moscow: few clouds`
```go
var ApiKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	var weatherRightNow, err = forecastle.CurrentWeatherById(524901, "metric", ApiKey, "us")
	if err != nil {
            log.Fatal(err)  
            return
        }

    fmt.Printf("Current weather in %s: %s", weatherRightNow.Name, weatherRightNow.Weather[0].Description)

}
```
## Documentation

See [repo wikia](https://github.com/h4desune/forecastle/wiki) for the documentation on how to install, run, use and etc.

## License
[Apache-2.0 License](http://www.apache.org/licenses/LICENSE-2.0)
