# Installation from source

1. Verify that you have Go 1.18+ installed

   ```shell
   $ go version
   ```

   If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).

2. Clone this repository

   ```sh
   $ git clone https://github.com/abc-inc/persephone.git
   $ cd persephone
   ```

3. Build and install

   #### Unix-like systems
   ```sh
   # installs to '/usr/local' by default; sudo may be required
   $ make install
   
   # or, install to a different location
   $ make install bindir=~/bin
   ```

   #### Windows
   ```pwsh
   # determine the version
   > git describe --tags
   # or
   > git rev-parse --short HEAD
   # build the `bin\persephone.exe` binary (replace VERSION with the actual version)
   > go build -trimpath -ldflags "-X main.version=VERSION" -o bin ./...
   ```
   There is no install step available on Windows.

4. Run `persephone version` to check if it worked.

   #### Windows
   Run `bin\persephone.exe version` to check if it worked.

## Cross-compiling binaries for different platforms

You can use any platform with Go installed to build a binary that is intended for another platform or CPU architecture.
This is achieved by setting environment variables such as `GOOS` and `GOARCH`.

For example, to compile the `persephone` binary for the 32-bit Raspberry Pi OS:

```sh
# on a Unix-like system:
$ GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 make
```

Run `go tool dist list` to list all supported values of `GOOS`/`GOARCH`.

Tip: to reduce the size of the resulting binary, you can use `GO_LDFLAGS="-s -w"`.
This omits symbol tables used for debugging.
See the list of [supported linker flags](https://golang.org/cmd/link/).

## Linking without Markdown viewer

persephone comes with a built-in Markdown renderer, which has multiple themes and is highly customizable.
If the binary size really matters, you can build persephone without renderer, which reduces the size by approx. 7 MB.
This can be achieved by setting `GOFLAGS="-tags=no_markdown"`.
Then it will simply write the Markdown text to stdout without formatting.
