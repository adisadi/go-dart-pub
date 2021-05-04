
# go-dart-pub
[![Docker Image](https://img.shields.io/docker/pulls/huiyiqun/pub_mirror.svg)](https://hub.docker.com/r/adisadi/go-dart-pub)

minimal private pub server written in go

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
$ docker pull adisadi/go-dart-pub
```

docker compose:

```bash
$ docker-compose up -d
```

## Usage

### Publish

Set the publish_to property in your pubspec to the `serving-url`

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

However, you might also consider:
* Using a reverse proxy for https, oauth etc...


Pub Custom Auth discussion:
    https://github.com/dart-lang/pub-dev/issues/4671




