RM ?= rm -rf
GOBUILD = go build
GOTEST = go test # -v

VARS=vars.mk
$(shell ./build_config ${VARS})
include ${VARS}

.PHONY: test clean build

build:
	@cd bin/ && ${GOBUILD} Main && cd ..

test:
	@cd src/FreqCtrl/ && ${GOTEST} && cd ../..

bench:
	@cd src/FreqCtrl/ && ${GOTEST} -bench=. && cd ../..

clean:
	${RM} ${VARS} bin/*
