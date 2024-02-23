scripts_dir := ./scripts

.PHONY: build

all: install setup run

build-core:
	$(scripts_dir)/build_core.sh

clean:
	rm -rf .build
	rm -rf .config
	rm -rf .next

dev-core:
	$(scripts_dir)/dev_core.sh

dev-web:
	$(scripts_dir)/dev_web.sh

install:
	$(scripts_dir)/install.sh

run:
	$(scripts_dir)/run.sh

setup:
	$(scripts_dir)/setup.sh "core"
	$(scripts_dir)/setup.sh "web"
