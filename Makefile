scripts_dir := ./scripts

.PHONY: build

all: install run

build:
	$(scripts_dir)/build.sh

install:
	$(scripts_dir)/install.sh

run:
	$(scripts_dir)/run.sh
