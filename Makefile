DATE=`date +%Y-%m-%dT%H:%M:%S`
ISO_DATE=${DATE}.000Z
LDFLAGS=-ldflags "-X main.date=${ISO_DATE}"

DIRS=`find . -type d -name "* - *"`

.PHONY: all
all: dirs ldflags

.PHONY: dirs
dirs:
	./build.sh

.PHONY: ldflags
ldflags:
	go build ${LDFLAGS} -o bin/ldflags ./ldflags
