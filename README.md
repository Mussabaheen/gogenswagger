# gogenswagger

gogenswagger generates signature and boiler plate code for the test cases using the swagger JSON file.

## Run Locally

Clone the project

```bash
  git clone https://github.com/Mussabaheen/gogenswagger.git
```

Go to the project directory

```bash
  cd gogenswagger
```

resolve `go.mod` file

```bash
  go mod tidy
```

run the project

```bash
  go run main.go
```

## Creating build

```bash
  cd gogenswagger
  make build
  ./build/gogenswagger [flags]
```

## Flags

```bash
  -json string
        Specify the Swagger JSON file path
  -language string
        Specify the language extension, currently supported languages js and go (default "go")
  -output string
        Specify the path for generated test packages (default "./generated")
```

## Supported Programming Languages

| Supported Language | Description             |
| ------------------ | ----------------------- |
| Go                 | `*.go` extension        |
| Javascript         | `*.js` extension (chai) |

## Contributing

Contributions are always welcome!

Feel free to report an issue, make improvement suggestions or open a pull request

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Support

For support, email at mussabaheen@gmail.com
