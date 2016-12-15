GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKER=docker
SODA=soda
GLIDE=glide

deps: 
	$(GLIDE) install

build:
	$(GOBUILD) -v -o gcon

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/gcon
	rm -rf gcon

test:
	$(GOTEST) -v ./grifts -race
	$(GOTEST) -v ./models -race
	$(GOTEST) -v ./actions -race

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
