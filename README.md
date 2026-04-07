# Locust
Consume your API's resources

NOTE: some of this functionality is still in progress. If you are reading this at the time of install, not all locust features may be available.

## Installation

### Linux

Run go install:

``` bash
go install https://github.com/Jonny-Burkholder/locust
```

Set the `LOCUSTPATH` environment variable on your system:

```bash
export LOCUSTPATH=usr/local/locust
```

Your locustpath can be any directory you choose. This is the default location locust will look for things like config and test collections. Alternatively, you can pass those values in as flags.

## Usage

The main function of locust is `swarm`. If you have your LOCUSTPATH set up, locust will first look there for a default test suite and config. If no config is set, locust will use the default config settings (print your config settings with `locust config`. If no test suite is selected, locust will error and print usage instructions.

For additional information, run `locust --help`, or for subcommand usage run `locust {subcommand} --help`

### API Testing

LOCUST exists to test your API in every conceivable way! Made primarily for load testing, LOCUST is also more than capable of running functional test suites, replacing cumbersome software with simple YAML suites. This is covered by the `locust test` command, which will run a test suite with the primary intention of completing the tests successfully, reporting errors when this does not occur. The main difference from `locust swarm` is that with the `swarm` command, the request swarm is explicitly intended to swell until the system breaks, reporting useful benchmarks and insights along the way.

### Benchmarking

In addition to sending requests and receiving responses, LOCUST collects benchmarking data to see where your API weak and strong points are. Couple this with robuts profiling to optimize your APIs.

Save runs to compare api versions

### Compare

Create and server html charts comparing different run

### Easy test suites

Instead of hand-writing test runs or using curl (nothing wrong with curl!), define simple YAML test suites and run them in one line from the terminal

## Looking for contributors!

Development of open-source software is hard, especially when we all have day jobs. We do it because we love free tech and sharing knowledge. If you like this project idea and would like to help, please reach out!

## Philosophy

I am a human, and I make software for other humans. I'm not against using AI in development, and I think it can be a powerful work accelerator. But for this project, I have chosen to write all of the code by hand. Despite having a copilot.md file, not a drop of this code is AI generated, and I'd like to keep it that way. Hand-writing code keeps me sharp, and it helps to ensure that some of the trickier code in this project is actually functional, and not just plausible. LLM-generated pull requests will be rejected. Feel free to use AI to help you understand the code and architecture, but please write the code yourself if you're looking to contribute

## Contributing

For more information on contributing, please refer to the [contribution guidelines](contribution_guidelines.md)
