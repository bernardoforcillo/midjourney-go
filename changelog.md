# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.1] - 2023-06-20

These are the changes are based on [@hongliang5316](https://github.com/hongliang5316/midjourney-go/) release.

### Added

- Complete the basic functions
- Support for `Imagine` api
- Support for `Upscale` api
- Support for `Variation` api
- Support for `Describe` api

## [0.1.0] - 2024-01-30

This code rewrite lays the groundwork for easier maintenance and reduces coupling between Discord and mid-journey functionality.

### Added

- Added stricted Discord API 

### Fixed

- Fixed Midjourney version

### Changed

- Changed module name from `hongliang5316/midjourney-go` to `bernardoforcillo/midjourney-go`
- Changed go version from `1.19` to `1.21`

## [0.1.1] - 2024-01-31

Code refator, cleanup and documentation.

### Changed

- Changed for deprecation of `ioutil` to `io`
- Changed `midjourney.New` to `midjourney.NewMidjourneyClient`

### Removed

- Removed reimplementation of `hongliang5316`'s `upscale` and `imagine` APIs from `midjourney-go` (`upscale` and `imagine` are now part of `MidjourneyClient` API).

## [0.1.2] - 2024-02-20

Code refator, cleanup and documentation.

### Changed

- added more time in the generation process for `Imagine` and `Upscale`

## [0.1.3] - 2024-02-20

Code improvement.

## [0.1.4] - 2024-02-20

Fix search message by prompt

### Fixed

- modified `SearchMesssageByPrompt` to search new version of the default prompt used by Midjourney

## [0.1.5] - 2024-02-20

Fix search messages by prompt

### Fixed

- modified `SearchMesssageByPrompt` into `SearchGeneratedMessage` and rewrited
- modified `SearchMesssageWithContent` into `SearchUploadMessage` and rewrited

## Changed

- changed waiting time
