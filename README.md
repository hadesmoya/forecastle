# forecastle

Forecastle is a go package for using OpenWeather API easily.

## Installation

Use the [go get](https://golang.org/cmd/go/) command to install forecastle.

```bash
go get github.com/h4desune/forecastle
```

## Usage

```go
import "github.com/h4desune/forecastle"

func main() {

forecastle.GetCurrentWeather(Moscow, metric, default, XXX)
// where XXX - is an actual APIKEY for openweathermap.org
// output is: It's 29.33° metric and a clear sky right now in the Moscow

} 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[Apache-2.0 License](http://www.apache.org/licenses/LICENSE-2.0)
