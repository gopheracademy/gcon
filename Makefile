GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOBUILD) -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKER=docker
DOCKERCOMPOSE=docker-compose
SODA=buffalo db
GLIDE=glide
BUFFALO=buffalo

deps: 
	$(GOCMD) get -u -t -v github.com/gobuffalo/buffalo/buffalo
	$(GLIDE) install
	$(GOGET) -u github.com/gobuffalo/buffalo  && $(GOINSTALL) github.com/gobuffalo/buffalo
	$(GOGET) -u github.com/markbates/pop      && $(GOINSTALL) github.com/markbates/pop

build:
	$(BUFFALO) build -o bin/gcon

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
	@echo "Make sure you've run 'make db-setup' before this"
	$(DOCKER) run --name=gophercon_db -d -p 5432:5432 -e POSTGRES_DB=gophercon_development postgres

db-setup: 
	$(DOCKERCOMPOSE) build
	$(DOCKER) run --name=gophercon_db -d -p 5432:5432 -e POSTGRES_DB=gophercon_development postgres
	sleep 6
	$(BUFFALO) db create -a
	$(BUFFALO) db migrate up
	$(DOCKER) ps | grep gophercon_db

db-down: 
	$(DOCKER) stop gophercon_db
	$(DOCKER) rm gophercon_db 

setup-dev: deps db-setup

teardown-dev: clean
	$(DOCKERCOMPOSE) down

run-dev: db-up
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
