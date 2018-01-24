NAME:=journeys
MAINPACKAGE:=github.com/hackathon/journeys
REALPKGDIR:=${GOPATH}/src/${MAINPACKAGE}
GINKGO_OPTS=-keepGoing -r -randomizeAllSpecs -randomizeSuites
SRCDIRS:=resources/
RESOURCE_PB=${wildcard resources/*/protocol/protocol.proto}
RESOURCE_PB_JSONS=${RESOURCE_PB:protocol.proto=protocol.pb_easyjson.go}
RESOURCE_PB_GO=${RESOURCE_PB:protocol.proto=protocol.pb.go}


#
# Development Compilation
#
$(NAME): ensure_deps protocol resourcefile/rodata.go
	@# go install performs an incremental compilation (a package will be recompiled if needed)
	go install -v ${MAINPACKAGE}


#
# Generate the Protocol
#
protocol: ensure_env ensure_deps protocols/protocols.go \
	${RESOURCE_PB_GO} \
	${RESOURCE_PB_JSONS}


resources/%/protocol/protocol.pb.go: resources/%/protocol/protocol.proto
	protoc --gofast_out=resources/$*/protocol/ -I ${GOPATH}/src -I ${GOPATH}/src/github.com/gogo/protobuf/protobuf -I resources/$*/protocol/ $<

resources/%/protocol/protocol.pb_easyjson.go: resources/%/protocol/protocol.pb.go
	cd "${REALPKGDIR}" && easyjson -all $<

#
# Linting
#
lint: protocol
	gometalinter --disable-all \
		-E gofmt -E vetshadow -E goconst \
		--tests \
		--exclude="[a-zA-Z\_]+_easyjson\.go" \
		--exclude="[a-zA-Z\_]+\.pb\.go" \
		--exclude="/resourcefile/rodata.go" \
		--vendor \
		--skip="protocol" \
		--deadline=10m  \
		./...
	! go list ./... | \
		fgrep -v -e /vendor/ -e /resourcefile | \
		xargs -L1 golint | \
		grep -v -E -e "[a-zA-Z\_]+\.pb\.go" -e "[a-zA-Z\_]+_easyjson\.go"

#
# Vetting
# lostcancel is disabled as it seems to trigger some false
# positives in Ginkgo tests. lostcancel is:
#   check for failure to call cancelation function returned by context.WithCancel
#
vet:
	cd "${REALPKGDIR}" && go tool vet -lostcancel=false ${SRCDIRS}


#
# Testing
#

ci-test: lint test

test: ensure_deps protocol resourcefile/rodata.go ensure_env ensure_test_env
	cd "${REALPKGDIR}" && ginkgo ${GINKGO_OPTS} -noisyPendings=false -race ${SRCDIRS}

profile: $(NAME)
	cd "${REALPKGDIR}" && ginkgo ${GINKGO_OPTS} -noisyPendings=false  -trace \
		-cpuprofile "cpu.profile" \
		-memprofile "mem.profile" \
		${SRCDIRS}

#
# Make our .banks files available inside our built executable
#
resourcefile/rodata.go: .banks/lock
	go-bindata -o $@ -pkg resourcefile -nomemcopy -prefix "${REALPKGDIR}" .banks
	gofmt -s -w $@


protocols/protocols.go: resources/*/protocol/protocol.proto
	rm -fr .protocols
	mkdir -p .protocols
	bash -c 'for x in resources/*/protocol/protocol.proto; do y=$${x:10}; ln -s ../$$x ".protocols/$${y%\/protocol\/protocol.proto}.proto"; done'
	go-bindata -o $@ -pkg protocols -nomemcopy -prefix .protocols .protocols
	rm -fr .protocols
	gofmt -s -w $@

#
# Vendor dependencies
#

vendor/.present: glide.lock
	mkdir -p vendor
	cd "${REALPKGDIR}" && glide --quiet install
	touch vendor/.present

ensure_deps: glide.lock vendor/.present

update_deps:
	cd "${REALPKGDIR}" && glide update && glide --quiet install
	touch vendor/.present


#
# Setup (used by `banks init`)
#
setup: ensure_env ensure_test_env update_deps



#
# Development dependencies
#
ensure_env:
	go get \
		github.com/mailru/easyjson/... \
		github.com/golang/protobuf/proto/... \
		github.com/golang/protobuf/protoc-gen-go/... \
		github.com/gogo/protobuf/proto/... \
		github.com/gogo/protobuf/protoc-gen-gogo/... \
		github.com/gogo/protobuf/gogoproto/... \
		github.com/gogo/protobuf/protoc-gen-gofast/... \
		github.com/jteeuwen/go-bindata/... \
		github.com/nats-io/gnatsd \
		github.com/autopilothq/lg \
		github.com/alecthomas/gometalinter
	gometalinter --install


update_env:
	go get -u \
		github.com/mailru/easyjson/... \
		github.com/golang/protobuf/proto/... \
		github.com/golang/protobuf/protoc-gen-go/... \
		github.com/gogo/protobuf/proto/... \
		github.com/gogo/protobuf/protoc-gen-gogo/... \
		github.com/gogo/protobuf/gogoproto/... \
		github.com/gogo/protobuf/protoc-gen-gofast/... \
		github.com/jteeuwen/go-bindata/... \
		github.com/nats-io/gnatsd \
		github.com/autopilothq/lg \
		github.com/alecthomas/gometalinter
	gometalinter --install


#
# Testing dependencies
#
ensure_test_env:
	go get \
		github.com/fzipp/gocyclo \
		github.com/onsi/ginkgo/ginkgo/... \
		github.com/onsi/ginkgo/extensions/table \
		github.com/onsi/gomega

update_test_env:
	go get -u \
		github.com/fzipp/gocyclo \
		github.com/onsi/ginkgo/ginkgo/... \
		github.com/onsi/ginkgo/extensions/table \
		github.com/onsi/gomega


#
# Release artifact compilation
#
ci-release: bin/${NAME}-linux-amd64

# Linker flags used
# -s  Omit the symbol table and debug information.
# -w  Omit the DWARF symbol table.
#
# https://golang.org/cmd/link/
#
bin/${NAME}-linux-amd64: ensure_deps protocol resourcefile/rodata.go
	GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w"  -o bin/${NAME}-linux-amd64 ${MAINPACKAGE}

#
# Cleaning
#
clean:
	go clean -i
	rm -f resourcefile/rodata.go
	rm -fr bin vendor


#
# Utility targets
#

.PHONY: ${NAME} protocol lint vet ci-test test profile ensure_deps update_deps setup ensure_env update_env ensure_test_env update_test_env ci-release
.SUBLIME_TARGETS: ${NAME} test
