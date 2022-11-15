build:
	rm -rf build && \
	mkdir build && \
	go build -o ./build/cliTools

clean:
	rm -rf build