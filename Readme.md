# Cgen | CLI Code Generator

The CLI Code Generator is a command-line interface (CLI) application designed to generate code based on templates provided in JSON format. It accepts a JSON configuration file and uses the Go standard library as a templating engine to generate code files.

## Installation

```bash
...conging soon
```

## Usage

The CLI Code Generator requires a JSON configuration file specifying the templates, destination paths, variables, and file renames. Here's how to use it:

```bash
./cgen -c <config_file.json>
```

Or using the long option:

```bash
./cgen --config <config_file.json>
```

If no configuration file is provided, the CLI Code Generator will look for a file named `.gen.config.json` in the current working directory.

### Sample Configuration File

```json
{
  "templates": {
    "path": "templates",
    "destination": "src",
    "variables": [
      {
        "name": "Name",
        "value": "John"
      },
      {
        "name": "ShowDetails",
        "value": "true"
      },
      {
        "name": "Age",
        "value": "30"
      },
      {
        "name": "Location",
        "value": "New York"
      }
    ],
    "fileRenames": [
      {
        "from": "_demo",
        "to": "Demo"
      }
    ]
  }
}
```

## Configuration Options

- **templates**: Object specifying the path to the templates directory, destination directory for generated code, variables to be substituted in templates, and file renaming rules.
    - **path**: Path to the directory containing template files.
    - **destination**: Path to the directory where generated code files will be saved.
    - **variables**: Array of objects specifying variable names and their values for substitution in templates.
    - **fileRenames**: Array of objects specifying rules for renaming files during generation.

## Template Syntax

The CLI Code Generator uses the Go standard library as a templating engine. Templates follow the Go text/template package syntax. Learn more about Go templates [here](https://golang.org/pkg/text/template/).
