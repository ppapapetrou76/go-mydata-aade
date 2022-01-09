SHELL := /bin/bash

# VARIABLES used
export GOBIN = $(shell pwd)/bin

# Include all makefiles used in the project
include build/Makefile.help
include build/Makefile.dev
include build/Makefile.deps
include build/Makefile.api
