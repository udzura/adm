VERBOSE_FLAG = $(if $(VERBOSE),-v)

build:
	go build $(VERBOSE_FLAG) -o adm.cmd ./cmd/adm
