# Go Gin Boilerplate
> A starter project with Golang, Gin and Gorm

[![Build Status][travis-image]][travis-url]
[![codebeat badge](https://codebeat.co/badges/ed248580-942c-4ffc-919f-d3681d28a799)](https://codebeat.co/projects/github-com-vsouza-go-gin-boilerplate)
[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_shield)


Golang Gin boilerplate with Gorm. Supports multiple configuration environments.

![](header.jpg)

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

__Install golang/dep__

```sh
brew install dep
```

then run:

```sh
dep ensure
```

## Development setup

```sh
cp config/development.example.yaml config/development.yaml
```

## Usage example

`curl http://localhost:8080/health`

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See [License](https://vsouza.mit-license.org) for more information.

[https://github.com/vsouza](https://github.com/vsouza)

[go-image]: https://img.shields.io/badge/Go--version-1.9-blue.svg
[go-url]: https://golang.org/doc/go1.9
[travis-image]: https://travis-ci.org/vsouza/go-gin-boilerplate.svg?branch=master
[travis-url]: https://travis-ci.org/vsouza/go-gin-boilerplate
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: https://vsouza.mit-license.org


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvsouza%2Fgo-gin-boilerplate?ref=badge_large)