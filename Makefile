scripts_dir := ./scripts

.PHONY: build

all: install setup run

build-core:
	$(scripts_dir)/build-core.sh

clean:
	rm -rf .build
	rm -rf .config

install:
	$(scripts_dir)/install.sh

run:
	$(scripts_dir)/run.sh

setup:
	$(scripts_dir)/setup.sh
