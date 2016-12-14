GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKER=docker
SODA=soda
GLIDE=glide

define GIT_ERROR
FATAL: Git (git) is required to download gcon dependencies.
endef

define HG_ERROR
FATAL: Mercurial (hg) is required to download gcon dependencies.
endef

define GLIDE_ERROR
FATAL: Glide (glide) is required to download gcon dependencies.
endef

ALL_PACKAGES = actions models grifts

deps: 
	$(GLIDE) install

build:
	$(GOBUILD) -v -o gcon

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/gcon
	rm -rf gcon

TEST_LIST = $(foreach pkg, $(ALL_PACKAGES), $(pkg)_test)
test: $(TEST_LIST)
$(TEST_LIST): %_test:
	if [ ! -d coverage ]; then mkdir coverage; fi 
	$(GOTEST) -v ./$* -cover -race -coverprofile=coverage/$(subst /,_,$*).coverprofile 

COVER_LIST = $(foreach pkg, $(ALL_PACKAGES), $(pkg)_cover)
cover: $(COVER_LIST)
$(COVER_LIST): %_cover:
	$(GOCMD) tool cover -html=coverage/$(subst /,_,$*).coverprofile -o coverage/$(subst /,_,$*).html

# check for git
git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

# check for mercurial
hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))

# check for glide
glide:
	$(if $(shell glide), , $(error $(GLIDE_ERROR)))
