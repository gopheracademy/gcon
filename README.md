# GopherCon 2017 [![Code Climate](https://codeclimate.com/github/gopheracademy/gcon/badges/gpa.svg)](https://codeclimate.com/github/gopheracademy/gcon) [![Build Status](https://travis-ci.org/gopheracademy/gcon.svg?branch=master)](https://travis-ci.org/gopheracademy/gcon)

## Setup

We're using [Glide](https://github.com/Masterminds/glide) for dependency management.  

### Installation 

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

Once Glide is installed, download the dependencies. 

```
glide install
```

Versions are managed in the `glide.yaml` file.  A `glide.lock` file is created

## Running

### Start Docker

```
docker run --name gcon -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
soda create -a
```

### Start Server

```
buffalo dev
```

## Testing

```
make test
```

### Start Docker -- See above.  

Tests REQUIRE a working database for now.

```
buffalo test
```
