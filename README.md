# Warehouse Rack CLI

Warehouse Rack CLI is a command-line application for managing a warehouse rack system. The application allows for interactive or file-based input to manage racks and check product statuses.

## Table of Contents

- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Installation](#installation)
- [Installation with Nix](#installation-with-nix)
- [Running the Application](#running-the-application)
  - [Interactive Mode](#interactive-mode)
  - [File Input Mode](#file-input-mode)
- [Running Tests](#running-tests)

## Project Structure

```bash
warehouse-rack-cli/
├── bin/                      # CLI executable scripts
│   ├── warehouse_rack        # Main executable
│   ├── setup                 # Build and setup script
│   └── run_functional_tests  # Script to run functional tests
├── cmd/                      # Entry point for Go main file
│   └── racking-system/       # Main Go application
├── functional_spec/          # Ruby functional test suite using PTY
│   ├── fixtures/             # contains file input
│   ├── spec/                 # RSpec tests
│   └── Gemfile               # Ruby dependencies
├── lib/                      # library codes
│   └── utils                 # utility functions
├── pkg/                      # Application code
│   ├── command/              # Command processor logic
│   └── rack/                 # Rack system logic
├── .envrc.example            # .envrc example file used by nix or direnv
├── .gitignore                # Git ignore
├── flake.lock                # Flake lock file
├── flake.nix                 # Flake configuration file for Nix
├── LICENSE                   # License
├── README.md                 # This file
└── go.mod                    # Go module dependencies
```

## Requirements

- Go 1.18+
- Bundler (for Ruby-based functional tests)
- Ruby 2.5+
- Optional: `make` (if using Makefile)
- Optional: `nix` (if using nix package manager)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/arifinoid/warehouse-rack-cli.git
   cd warehouse-rack-cli
   ```

2. Install Go dependencies and build the project:

   ```bash
   bin/setup
   ```

3. (Optional) Install Ruby dependencies for running functional tests:

   ```bash
   cd functional_spec
   bundle install
   ```

## Installation with Nix

```bash
cp .envrc.example .envrc
direnv allow
```

## Running the Application

You can run the application in two modes: interactive mode or file input mode.

### Interactive Mode

Run the application without any arguments to start in interactive mode:

```bash
bin/warehouse_rack
```

You can then enter commands interactively. For example:

```bash
$ create_rack 6
Created a warehouse rack with 6 slots

$ rack ZG11AQA 2024-02-28
Allocated slot number: 1

$ status
Slot No.    SKU No.     Exp Date
1           ZG11AQA     2024-02-28

$ exit
```

### File Input Mode

To process a batch of commands from a file, pass the file path as an argument:

```bash
bin/warehouse_rack path/to/input_file.txt
```

For example:

```bash
bin/warehouse_rack functional_spec/fixtures/file_input.txt
```

## Running Tests

There are two types of tests: Go unit tests and functional tests using Ruby PTY.

### Unit Tests

Run Go unit tests with the following command:

```bash
go test ./pkg/rack/
```

### Functional Tests

To run functional tests using the Ruby PTY module:

1. Make sure you're in the `functional_spec` directory.
2. Run the following command:

   ```bash
   PATH=$PATH:../bin bundle exec rake spec:functional
   ```

You can also run a specific functional test:

```bash
PATH=$PATH:../bin bundle exec rspec spec/warehouse_rack_spec.rb --tag sample
```
