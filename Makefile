# See LICENSE for details
# makefile for goric

.PHONY: all build clean format format-check lint run
.DEFAULT: all

include config.mk

PKGFILES = \
	cmd\
	config.mk\
	go.mod\
	hostname.go\
	LICENSE\
	Makefile\
	nslook.go\
	README.md\
	tcp\
	udp

BUILD_DIR = target
DIST_DIR = $(BUILD_DIR)/dist
DIST_BASE_DIR = goric-$(VERSION)
BIN=goric

all: $(BIN)
	@echo "goric built"

clean:
	@echo cleaning
	@rm -rf $(BIN) $(BUILD_DIR)

dist:
	mkdir -p $(DIST_DIR)/$(DIST_BASE_DIR)
	cp -r $(PKGFILES) $(DIST_DIR)/$(DIST_BASE_DIR)
	cd $(DIST_DIR); \
	tar -cz  -f $(DIST_BASE_DIR).tar.gz $(DIST_BASE_DIR); \
	zip -r $(DIST_BASE_DIR).zip $(DIST_BASE_DIR)

format:
	gofumpt -w .

format-check:
	gofumpt -l .

install: $(BIN)
	@echo installing executable file to $(PREFIX)/bin
	@mkdir -p $(PREFIX)/bin
	@cp -f $(BIN) $(PREFIX)/bin
	@chmod 755 $(PREFIX)/bin/goric

lint:
	golangci-lint run .

options:
	@echo "goric compilation flags"
	@echo "VERSION    = $(VERSION)"
	@echo "BIN        = $(BIN)"

uninstall:
	@echo removing executable file from $(PREFIX)/bin
	@rm -f $(PREFIX)/bin/ric

$(BIN):
	go build -o $(BIN) ./cmd/goric
