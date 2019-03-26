clean:
	rm clean.txt
	rm messy.txt

clean.txt:
	go run --tags='clean' . |tee clean.txt

messy.txt:
	go run --tags='messy' . |tee messy.txt

compare: clean messy.txt clean.txt
ifeq (, $(shell which benchstat))
$(error "No benchstat for install: go get golang.org/x/perf/cmd/benchstat")
endif
	benchstat messy.txt clean.txt

