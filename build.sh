#!/usr/bin/env bash
RUN_NAME="KitexFfmpeg"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -ldflags '-s -w -L /usr/lib -linkmode "external" -extldflags "-static"' -o output/bin/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
