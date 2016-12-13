# GopherCon 2017 [![Code Climate](https://codeclimate.com/github/gopheracademy/gcon/badges/gpa.svg)](https://codeclimate.com/github/gopheracademy/gcon) [![Build Status](https://travis-ci.org/gopheracademy/gcon.svg?branch=master)](https://travis-ci.org/gopheracademy/gcon)

# Running

## Start Docker

	 docker run --name gcon -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	 soda create -a

## Start Server

	buffalo dev

# Testing

## Start Docker -- See above.  

Tests REQUIRE a working database for now.

	buffalo test
