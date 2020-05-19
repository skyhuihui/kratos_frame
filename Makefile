.PHONY: all build  run  clean

all:
	@echo "kratos_frame"

build:
	cd cmd && go build

run:
	cd cmd && nohup ./cmd -conf ../configs &

clean:
	@rm -rf cmd/cmd
