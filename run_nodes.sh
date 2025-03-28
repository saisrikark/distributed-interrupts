#!/bin/bash

set -e

if [ ! -f "bin/noderunner" ]; then
    make noderunner
fi

./bin/noderunner 