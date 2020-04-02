#!/usr/bin/env bash

GIT_REVISION=$(git rev-parse --short --verify HEAD)
TIME=$(date -u +%Y%m%d.%H%M%S)
VERSION=1.0.${GIT_REVISION}.${TIME}

build_artifacts () {
  local os=$1
  local arch=$2
  GOOS=$os GOARCH=$arch go build -ldflags "-X main.versionString=${VERSION}" -o keells-check ./main.go
  file keells-check
  tar -czvf keells-check-${os}-x64.tar.gz keells-check
  rm keells-check
}

build_artifacts linux amd64
build_artifacts darwin amd64
build_artifacts windows amd64