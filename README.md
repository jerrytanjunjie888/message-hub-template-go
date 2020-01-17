# Message Hub Template Module

This module converts the structure of incoming parameters from the iCheck API which is developed/managed in France.
The adapter receives the incoming requests, restructures them to message Hub's api style/structure. It then sends them to
another module in message hub to prepare it for message sending.
s
## Architecture

All our modules follow the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). We strongly ensure that all layers of the application are highly cohesive and yet very independent of each other. Layers interact using interfaces.
We implement the **Service**, **Repository**, **Model** and **Delivery** layers.

To get the high level understanding of Messege hub, follow this. [Message-Hub](https://docs.google.com/presentation/d/1gHSOl33SVmchPIy6YesP9oEAcMZy2OVsWFT_TgGupmM/edit?usp=sharing)

### Built withx

* [Go](https://golang.org/) Programming Language

#### Essential Packages

* [Gin-Gonic](https://gin-gonic.com/) - Web Routing
* [Uber Zap](https://github.com/uber-go/zap) - Logging
* [LumberJack](https://github.com/natefinch/lumberjack) - Log Rotation

### Getting Started

#### Prerequisites

* Start by cloning this project

```bash
git clone -b develop git@github.com:dktunited/message-hub-adapter.git
```

* Make sure you have Go installed. Check out this [guide](https://golang.org/doc/install) to installing Go. For macOS users you can run
this simple command to install:

```bash
brew install go
```

* Add the Go path to your environment variable by running this:

```bash
export PATH=$PATH:/usr/local/go/bin
```

* Enable Go Module. *cd into message-hub-adapter*

```bash
export GO111MODULE=on
```

* Install the essential packages:

 	* Gin-Gonic

		```bash
		go get -u github.com/gin-gonic/gin
		```

  * Uber Zap

	```bash
	go get -u go.uber.org/zap
	```

  * LumberJack

	```bash
	go get -u github.com/natefinch/lumberjack
	```

* Download/verify go packages

```bash
go mod download && go mod tidy && go mod verify
```

To successfully launch this module, you need to do this:

* Set a path on your local computer to save the log files. Example:  ``/data/messaghub/adapter/``

* Change all environment variables to your preferred choice. See LumberJack's [documentation](https://github.com/natefinch/lumberjack).

``` go
logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
logMaxBackups, _ := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
logMaxAge, _ := strconv.Atoi(os.Getenv("LOG_MAX_Age"))
logCompress, _ := strconv.ParseBool(os.Getenv("LOG_COMPRESS"))

hook := &lumberjack.Logger{
 Filename:   os.Getenv("LOG_DIR") + os.Getenv("LOG_FILENAME"),
 MaxSize:    logMaxSize,
 MaxBackups: logMaxBackups,
 MaxAge:     logMaxAge,
 Compress:   logCompress,
}
```

* Make sure you map your local log path to the log path in the docker container.
  Do this in your ``docker-compose`` file:

```yml
  volumes:
    - /YOUR_LOCAL_LOG_PATH/:/tmp/
 ```

Run the following commands to launch application

##### Run on local

```bash
go run main.go
```

##### Run with Docker

Run ``up.sh`` to start up and ``down.sh`` to shutdown

```bash
up.sh
```

```bash
down.sh
```

## Test

### Local

For local test, check out the **Request.rest** file in the `/test folder`.
You can copy the examples in there and run using curl or Postman.

### Pre-production

On the other hand, you can directly test with our [preproduction](https://api-portal.preprod.subsidia.org/#!/apis/8b2fb3af-7c9a-4270-afb3-af7c9ab2703f/pages/f3863dc8-8758-4a6f-863d-c887583a6fcf) environment on Api-Portal

## Roadmap

#### Upcoming

###### Q1 2020

###### Q2 2020

#### Changelog

## Support
For all suggestions, bug reports and feature requests please feel free
to raise an **issue** or check out this [Support document](https://docs.google.com/document/d/1wOpphMC9qt3U6IvbSlUmu68AGelhLQhE9XxhigyB0iY/edit?usp=sharing).
