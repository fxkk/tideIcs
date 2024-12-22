# TideIcs

TideIcs is a Go project to transform tide information into iCalendar format.
Tide Data can be obtained from [bsh](https://www.bsh.de/DE/DATEN/Vorhersagen/Gezeiten/gezeiten_node.html) in txt format.

I just created this project for myself, if it can be useful for someone, even better.

## Installation

To install TideIcs, use `go get`:

```sh
go install github.com/fxkk/tideIcs
```

or

```sh
version=$(curl https://api.github.com/repos/fxkk/tideIcs/releases/latest -s | jq .name -r)
curl -OL  "https://github.com/fxkk/tideIcs/releases/download/${version}/tideIcs_$(uname)_$(uname -m).tar.gz"
tar xzf tideIcs_$(uname)_$(uname -m).tar.gz
chmod u+x tideIcs
./tideTcs --help
```

## Usage

To use TideIcs, run the following command:

```sh
tideIcs tide_data.txt > tide.ics
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.