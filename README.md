# yar

A project by [@draganm](https://github.com/draganm)

## Description

`yar` (YAML Archiver) is a command-line tool that archives directory contents into a YAML file. It recursively walks through a directory and stores all file contents in a YAML format, where file paths are keys and their contents are values.

## Installation

```bash
go install github.com/draganm/yar@latest
```

## Usage

```bash
# Archive a directory into a YAML file
yar --output output.yaml <directory>

# Using environment variable for output
export YAR_OUTPUT=output.yaml
yar <directory>
```

### Flags
- `--output, -o`: Output file path (required)
  - Can also be set via `YAR_OUTPUT` environment variable

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.txt) file for details.

## Author

- [@draganm](https://github.com/draganm)
