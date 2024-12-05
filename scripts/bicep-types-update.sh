#!/bin/bash

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOTDIR="$MYDIR/.."

main() {
    bicep_dir="$ROOTDIR/embed/bicep-types-az/generated"
    echo "bicep_dir: $bicep_dir"
    target_dir="$ROOTDIR/embed/generated"
    echo "target_dir: $target_dir"
    echo "removing all exist type files..."
    rm -r $target_dir
    echo "done"
    echo "copying new type files..."
    cp -r $bicep_dir $target_dir
    echo "done"
    cd $target_dir
    echo "removing all .md and .out files..."
    find . -name "*.md" -type f -delete
    find . -name "*.out" -type f -delete
}

main "$@"