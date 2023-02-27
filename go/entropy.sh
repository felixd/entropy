#!/bin/bash
# Author: Paweł 'felixd' Wojciechowski - FlameIT - Immersion Cooling - https://flameit.io
FILE=random.bin

echo "Generating random numbers."
echo "go run entropy.go -n NUMBER_OF_BYTES [default: 8*1000000]"
echo "16 bytes should ideally give: 16 * 8 = 128 bit entropy"
echo "32 bytes should ideally give: 32 * 8 = 256 bit entropy"

echo "Testing 256 bits / 32 bytes"

go run entropy.go -n 32

ls -al $FILE
echo "od -t x1 $FILE"
od -t x1 $FILE

echo "od -x $FILE"
od -x $FILE

echo "Dieharder testing:"
dieharder -f $FILE -a

echo "Done"
