all: build

clean:
	rm -rf pkg bin

build: clean
	gb build all

install: build

update:
