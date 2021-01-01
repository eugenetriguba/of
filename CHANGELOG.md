# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.1] - 2021-01-01

- Refactor todo `Send()` to use github.com/jordan-wright/email instead of gomail.v2.
- Change package into a go module

## [0.2.0] - 2020-01-01

### Added

- `of add` command to add a task to your omnifocus inbox.
- `of version` command to check the version of the Omnifocus CLI.
- `of config` command to manage configuration settings.
- `of config output` command to add ability to output the config to stdout.
- More informative error messages

## Changed

- Configuration is now handled in `~/.of/` folder instead of locally in the package.

## [0.1.0] - 2019-12-28

- Initial Release
