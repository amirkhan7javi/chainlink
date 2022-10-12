#!/usr/bin/env bash

set -ex

optimize_runs="$1"
solpath="$2"
solcoptions=("--optimize" "--optimize-runs" "$optimize_runs" "--metadata-hash" "none")

basefilename="$(basename "$solpath" .sol)"
pkgname="$(echo $basefilename | tr '[:upper:]' '[:lower:]')"

here="$(dirname $0)"
pkgdir="${here}/${pkgname}"
mkdir -p "$pkgdir"
outpath="${pkgdir}/${pkgname}.go"
abi="${pkgdir}/${basefilename}.abi"
bin="${pkgdir}/${basefilename}.bin"

SOLC_VERSION="0.8.0"

solc-select use $SOLC_VERSION
solc --version | grep "$SOLC_VERSION" || ( echo "You need solc version $SOLC_VERSION" && exit 1 )

# FIXME: solc seems to find and compile every .sol file in this path, so invoking this once for every file produces n*3 artifacts
solc "$solpath" ${solcoptions[@]} --abi --bin --combined-json bin,bin-runtime,srcmap-runtime --overwrite -o "$(dirname $outpath)"

go run wrap.go "$abi" "$bin" "$basefilename" "$pkgname"
