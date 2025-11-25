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

### Architecture
ScriptureBot is built as a 5 layer service:
1. Web App (GET)
2. Incoming Translation Layer from Platform specific properties
3. Logic Layer
4. Outgoing Translation Layer to Platform specific properties
5. Web App (POST)

The Translation Layer is implemented in [BotPlatform](http://github.com/julwrites/BotPlatform), which abstracts all the translation tasks from the Logic layer. 

Additionally there is a [BotSecrets](http://github.com/julwrites/BotSecrets) integration with the WebApp layer which provides all sensitive data to the bot on a as-needed basis.

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

### Build and Test

On a fresh repository, run `go mod init` and `go mod tidy` to get all the necessary go modules for runtime

To test, run `go test github.com/julwrites/ScriptureBot/pkg/<module>`, e.g `go test github.com/julwrites/ScriptureBot/pkg/app`

### CI/CD Pipeline
This repository uses go module to manage dependencies, and is hosted on gcloud cloud run. 

As such it requires [gcloud CLI](https://cloud.google.com/sdk/docs/quickstart) to package the Dockerfile

The artifact repository is set to `us-central1`

As such the docker container can be built using the following command
`docker build -f Dockerfile -t us-central1-docker.pkg.dev/${GCLOUD_PROJECT_ID}/scripturebot/root:latest .`

And then uploaded using

`docker push us-central1-docker.pkg.dev/${GCLOUD_PROJECT_ID}/scripturebot/root:latest`

And finally deployed using
`gcloud run deploy scripturebot --image us-central1-docker.pkg.dev/${GCLOUD_PROJECT_ID}/scripturebot/root:latest --region us-central1`
