GOCMD=go
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOCMD) build -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKERCOMPOSE=docker-compose
SODA=buffalo db
GLIDE=glide
BUFFALO=buffalo

deps: 
	$(GOCMD) get -u -t -v github.com/gobuffalo/buffalo/buffalo
	$(GLIDE) install

build:
	$(GOBUILD) -v -o gcon

buildprod:
	$(GOBUILDPROD) -v -o gcon

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/gcon
	rm -rf gcon

test:
	$(GOTEST) -v ./grifts -race
	$(GOTEST) -v ./models -race
	$(GOTEST) -v ./actions -race

db-up: 
	docker run --name=gophercon_db -d -p 5432:5432 -e POSTGRES_DB=gophercon_development postgres
	sleep 10
	$(SODA) create 
	$(SODA) migrate up
	docker ps | grep gophercon_db

db-down: 
	docker stop gophercon_db
	docker rm gophercon_db 

setup-dev: deps
	$(DOCKERCOMPOSE) build
	$(DOCKERCOMPOSE) up -d
	$(SODA) create -a
	docker ps | grep gcon_db

teardown-dev: clean
	$(DOCKERCOMPOSE) down

run-dev: 
	$(BUFFALO) dev

define GIT_ERROR
FATAL: Git (git) is required to download gcon dependencies.
endef

define HG_ERROR
FATAL: Mercurial (hg) is required to download gcon dependencies.
endef

define GLIDE_ERROR
FATAL: Glide (glide) is required to download gcon dependencies.
endef

# check for git
git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

# check for mercurial
hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))

# check for glide
glide:
	$(if $(shell glide), , $(error $(GLIDE_ERROR)))
