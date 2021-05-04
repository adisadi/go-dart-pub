
# go-dart-pub
minimal private pub server written in go

## How it works
`go-dart-pub` is a bare-minimum implementation of the Pub server API (https://github.com/dart-lang/pub/blob/master/doc/repository-spec-v2.md).
It stores the packages in Filsystem

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
$ docker pull todo
```

## Usage

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
* Using a reverse proxy


Pub Custom Auth discussion:
    https://github.com/dart-lang/pub-dev/issues/4671




