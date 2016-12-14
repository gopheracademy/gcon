# GopherCon 2017 [![Code Climate](https://codeclimate.com/github/gopheracademy/gcon/badges/gpa.svg)](https://codeclimate.com/github/gopheracademy/gcon) [![Build Status](https://travis-ci.org/gopheracademy/gcon.svg?branch=master)](https://travis-ci.org/gopheracademy/gcon)

# Running

## Start Database Server (Postgres)

	 docker-compose up -d

## Load database tables

You only need to run this command when you're starting up the project for the very first time.

	 soda create -a

## Start Server

	buffalo dev

## Shutdown Postgres after development

	docker-compose stop

# Testing

## Start Docker -- See above.  

Tests REQUIRE a working database for now.

	buffalo test
