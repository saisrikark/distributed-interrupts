#!/bin/bash

set -e

if [ ! -f "bin/noderunner" ]; then
    echo "Building noderunner..."
    make noderunner
fi

./bin/noderunner 