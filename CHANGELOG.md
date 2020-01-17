# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
  * Refactor todo `Send()` to use github.com/jordan-wright/email instead of gomail.v2.
  * 

## [0.2.0] - 2020-01-01
### Added
 - Cobra to manage CLI interactions. Application is reorganized accordingly.
 - `fs` package for file system related helper functions: `CloseFile()`, `DirExists()`, and `FileExists()`.
 - Documentation for `todo` package and more informative error messages.
 - Documentation for `configuration` package.
 - `GetConfigDirPath()` and `GetConfigFilePath()` as helpers for those paths in the `configuration` package.
 - `Init()` in `configuration` package to ensure that the configuration folder and file are present (and created if not).
 - `Output()` in `configuration` package to be able to output the configuration file to stdout.
 - `of version` command to check the version of the Omnifocus CLI.
 - `of config` command to manage configuration settings.
 - `of config output` command to add ability to output the config to stdout.
 - `of add` command to add a task to your omnifocus inbox.
### Changed
 - Configuration is now handled in `~/.of/` folder instead
   of locally in the package.

## [0.1.0] - 2019-12-28
### Added
 - Added conf.json configuration file to hold mail drop email, gmail email, and gmail password.
 - Added `of` command that parses `-email`, `-note`, and `-attachment` flags.
 - Added `configuration` type that represents the fields in the conf.json file. This type
   can parse the configuration file and write out to it.
 - Added `todo` type that holds information related to what we can send to 
   omnifocus to create todos. This type has the ability to send itself as an email.
 - Added `of` command ability to send a todo.