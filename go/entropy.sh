#!/bin/bash
# Author: Paweł 'felixd' Wojciechowski - FlameIT - Immersion Cooling - https://flameit.io

echo "Generating random numbers"
go run entropy.go

echo "Testing with Dieharder"
dieharder -f random.bin -a
