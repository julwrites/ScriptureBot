## Scripture Bot

![status: active](https://img.shields.io/badge/status-active-green.svg)

This Telegram bot hopes to make the Bible more accessible, and hopefully to give a timely answer to those looking for it. 

### Feedback
Star this repo if you found it useful. Use the github issue tracker to give
feedback on this repo.

## Licensing
See [LICENSE](LICENSE)

## Author
Hi, I'm [Julwrites](http://www.tehj.io)

## Code Guidelines

### Code
We are using Go 1.12 for this version of the framework.

Naming Convention:
* Variables should be named using camelCase.
* Methods should be named using underscore_case.
* Classes should be named using PascalCase.
* Packages should be named using underscore_case, in keeping with Python STL.
* Constants should be named using CAPITALCASE

This keeps the entities visually separate, and syntax clean.

As much as possible, each file should contain one of 3 things:
* A class and related methods
* A set of utility methods
* Business logic/End point logic

This is intended to facilitate separation of responsibility for loose coupling. 
