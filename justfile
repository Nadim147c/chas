name := "chas"
bin := "build" / name
tag := `git describe --long --tags --abbrev=7 | sed 's/^v//;s/\([^-]*-g\)/r\1/;s/-/./g'`
version := env("VERSION", tag)
user-local := `echo ~/.local`
prefix := absolute_path(env("PREFIX", user-local))

build-install:
    @just build
    @just install

install:
    install -Dm755 {{ bin }} "{{ prefix }}/bin/{{ name }}"

build:
    # Build for default os and arch
    @just compile "" ""

compile os arch:
    env \
      CGO_ENABLED=0 \
      GOOS="{{ os }}" \
      GOARCH="{{ arch }}" \
      go build -trimpath -ldflags '-s -w -X main.Version={{ version }}' -o {{ bin }}

version:
    @echo {{ version }}

bench:
    go test -bench=. -run=^$ ./...
