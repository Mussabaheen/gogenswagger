# gogenswagger

gogenswagger generates signature and boiler plate code for the test cases using the swagger JSON file.

## WIP

currently, work is done to integrate other langauges and improve the test coverage for the project.

## Usage

```bash
make build
./build/gogenswagger [swagger json path]
```

### Running the linter

```bash
make lint
```

## Supported Programming Languages

| Supported Language | Description             |
| ------------------ | ----------------------- |
| Go                 | `*.go` extension        |
| Javascript         | `*.js` extension (chai) |

## Contributing

Feel free to report an issue, make improvement suggestions or open a pull request

## License

[MIT](https://choosealicense.com/licenses/mit/)
