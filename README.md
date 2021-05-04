
# go-dart-pub
![Docker Pulls](https://img.shields.io/docker/pulls/adisadi1000/go-dart-pub?style=flat-square)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/adisadi1000/go-dart-pub?style=flat-square)
![Docker Automated build](https://img.shields.io/docker/automated/adisadi1000/go-dart-pub?style=flat-square)

minimal private pub server written in go

the indention of this project is a minimal microservice private pub server.
Just a a package repo, no auth no blingbling.

if you want tls, put it behind a reverse proxy.
if you want auth, solve it in a proxy.
if you want an other filesystem, solve it on infrastructure level (aws,blob storage, ..what ever)
if you want package statistics, solve it ...i think you got it, not here!.

## How it works
`go-dart-pub` is a bare-minimum implementation of the [Pub server API](https://github.com/dart-lang/pub/blob/master/doc/repository-spec-v2.md).
Filesystem is used as Storage.

For example, if you set up `testdriver` 'publish_to: http://localhost:8080' , then
you can include `package:testdriver` in your own packages, provided that you set up
Pub correctly (either via `PUB_HOSTED_URL`, or explicit `hosted` dependencies.)

```bash
export PUB_HOSTED_URL=http:localhost:8080
```

And then commands like `pub get`,`pub publish` and `pub upgrade` would work, seamlessly.

## Installation

as a docker image by:

```bash
$ docker pull adisadi1000/go-dart-pub
```

docker compose:

```bash
# see docker-compose.yml in repo
$ docker-compose up -d
```

## Usage

### Publish

Set the `publish_to` property in your pubspec to the `serving-url`

```yml
name: testdrive
description: A starting point for Dart libraries or applications.
publish_to: http://localhost:8080
version: 1.0.1
homepage: https://www.example.com
...
```

```bash
$ dart pub publish
```

### Consume

Setting the environment `PUB_HOSTED_URL` to the `serving-url`, then both `pub`
and `flutter` will download packages from pub Repo.

```bash
$ export PUB_HOSTED_URL="http://localhost:8080"
$ pub get  # downloaded from http://localhost:8080
$ flutter packages get  # downloaded from http://localhost:8080
```

or in your pubspec

```yml
dependencies:
  testdriver:
    hosted:
      name: testdriver
      url: http://localhost:8080
    version: ^1.0.0
```

## However, you might also consider:
* Using a reverse proxy for https, oauth etc...


Pub Custom Auth discussion:
    https://github.com/dart-lang/pub-dev/issues/4671




