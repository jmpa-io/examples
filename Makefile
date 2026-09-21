
# Default PROJECT, if not given by another Makefile.
ifndef PROJECT
PROJECT=examples
endif

# Targets.
build-image-%: cmd/%/Dockerfile
	docker build -f $< .

---: ## ---

# No race detector — OOMs on this machine.
GO_TEST_ARGS := -short -coverprofile=dist/coverage.txt -covermode=atomic -vet=off

# No src/ directory in this repo.
PYTHON_COV_ARGS :=

# Includes the common Makefile.
# NOTE: this recursively goes back and finds the `.git` directory and assumes
# this is the root of the project. This could have issues when this assumtion
# is incorrect.
include $(shell while [[ ! -d .git ]]; do cd ..; done; pwd)/Makefile.common.mk

