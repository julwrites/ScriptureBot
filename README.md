## Biblica Bot

![status: active](https://img.shields.io/badge/status-active-green.svg)

This Telegram bot hopes to make the Bible more accessible, and hopefully to give a timely answer to those looking for it. 

### Installing Libraries
Add new libraries either in app.yaml or add them to requirements.txt (Note that adding them to requirements.txt won't automatically import the libraries! You need to pip install them into a \lib folder still)

### Feedback
Star this repo if you found it useful. Use the github issue tracker to give
feedback on this repo.

## Licensing
See [LICENSE](LICENSE)

## Author
Hi, I'm [Julwrites](http://www.tehj.org)

## Code Guidelines

### Modules
This framework aims to implement a lightweight and reusable set of packages. 

The framework may be treated as a tree of packages, each of which contains at least one of the following core blocks:
* Utilities - Utility functions, which are the main logic
* Actions   - User-triggered functionalities, including states and interactions
* Hooks     - System-triggered functionalities which occur by cron job
e.g. package admin can contain admin_classes, admin_utils, admin_actions, admin_hooks

Specialized functionality can be implemented within each package however the implementer likes, but external access to the package should be limited to the core files only. This means, for example, that although the admin package may contain other modules, external users should not assume these modules exist. 

### Code
We are using Python27 for this version of the framework.

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
