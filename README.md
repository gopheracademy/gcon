# GopherCon 2017 [![Code Climate](https://codeclimate.com/github/gopheracademy/gcon/badges/gpa.svg)](https://codeclimate.com/github/gopheracademy/gcon) [![Build Status](https://travis-ci.org/gopheracademy/gcon.svg?branch=master)](https://travis-ci.org/gopheracademy/gcon) [![Go Report Card](https://goreportcard.com/badge/github.com/gopheracademy/gcon)](https://goreportcard.com/report/github.com/gopheracademy/gcon)

## Getting Started

The GopherCon 2017 website code requires a number of components to successfully build, deploy, and run.  For dependency management, we're using [Glide](https://github.com/Masterminds/glide).

* UNIX/Linux 

```
curl https://glide.sh/get | sh
```

* MacOSX 

```
brew install glide
```

* Ubuntu 

```
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
```

Versions are managed in the `glide.yaml` file and a subsequent `glide.lock` file is created.

Once Glide is installed we need to make sure that we have the PostgreSQL image downloaded, have the container built, schema loaded, and dependencies installed. To accomplish this, run the below command.
```
make setup-dev
```

Make sure that the project is building successfully.
```
make test
```

Run the project.
```
make run-dev
```

## Testing

```
make test
```

## Run Buffalo Directly

```
buffalo dev
```

## Run Buffalo Tests

```
buffalo test
```


## Create Admin User

```
buffalo task admin 
```
