# tiamat

Bring chaos to your API

NOTE: some of this functionality is still in progress. If you are reading this at the time of install, not all tiamat features may be available.

## Installation

### Linux

Run go install:

``` bash
go install https://github.com/Jonny-Burkholder/tiamat
```

Set the `TIAMATPATH` environment variable on your system:

```bash
export TIAMATPATH=usr/local/tiamat
```

Your tiamatpath can be any directory you choose. This is the default location tiamat will look for things like config and test collections. Alternatively, you can pass those values in as flags.

## Usage

The main function of tiamat is `consume`. If you have your tiamatPATH set up, tiamat will first look there for a default test suite and config. If no config is set, tiamat will use the default config settings (print your config settings with `tiamat config`). If no test suite is selected, tiamat will error and print usage instructions.

For additional information, run `tiamat --help`, or for subcommand usage run `tiamat {subcommand} --help`

### API Testing

tiamat exists to test your API in every conceivable way! Made primarily for load testing, tiamat is also more than capable of running functional test suites, replacing cumbersome software with simple YAML suites. This is covered by the `tiamat test` command, which will run a test suite with the primary intention of completing the tests successfully, reporting errors when this does not occur. The main difference from `tiamat consume` is that with the `consume` command, the request consume is explicitly intended to swell until the system breaks, reporting useful benchmarks and insights along the way.

### Benchmarking

In addition to sending requests and receiving responses, tiamat collects benchmarking data to see where your API weak and strong points are. Couple this with robuts profiling to optimize your APIs.

Save runs to compare api versions

### Compare

Create and server html charts comparing different runs

### Easy test suites

Instead of hand-writing test runs or using curl (nothing wrong with curl!), define simple YAML test suites and run them in one line from the terminal

## Suggested uses

In addition to the face value built-in functionality, there are many benefits to running an army of requests against your API:

- Profiling: Run your application with profiling enabled, and use tiamat to investigate where resources are being used.
- Auto-scaling: in a cloud-native environment, use tiamat to consume resources and ensure scaling is working as expected. Use exponential requests to cause usage to ramp up, then back down, over a specified time period, and measure your scaling response.

## Looking for contributors!

Development of open-source software is hard, especially when we all have day jobs. We do it because we love free tech and sharing knowledge. If you like this project idea and would like to help, please reach out!

## Philosophy

I am a human, and I make software for other humans. I'm not against using AI in development, and I think it can be a powerful work accelerator. But for this project, I have chosen to write all of the code by hand. Despite having a copilot.md file, not a drop of this code is AI generated, and I'd like to keep it that way. Hand-writing code keeps me sharp, and it helps to ensure that some of the trickier code in this project is actually functional, and not just plausible. LLM-generated pull requests will be rejected. Feel free to use AI to help you understand the code and architecture, but please write the code yourself if you're looking to contribute

## Contributing

For more information on contributing, please refer to the [contribution guidelines](docs/CONTRIBUTING.md.md)
