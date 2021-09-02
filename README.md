# Performance check for Golang Switch/Map

## How to run the present

1. Install the present tools from Golang

	go install golang.org/x/tools/present

2. Start the present in the current directory

	present

3. Open your browser and go to http://127.0.0.1:3999

## How to generate different size of elements

	$ cd plugs
	$ ELEMENTS=<size-of-elements> make clean all

**Note** you might encounter into a problem when using a too large size for the elements.

## How to check the performance?

	$ cd internal/plugs
	$ go test -v -bench=. -benchmem -benchtime 5s .

## Next steps

- [ ] Disassemble the obj files
  - [ ] See what's the hash function in Golang
  - [ ] See switch optimization from Golang
- [ ] Check the implementation from Golang
