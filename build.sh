#!/usr/bin/env bash
# build.sh
# Description: Build script for micro tools
# Author: BlueFalconHD

# Check if dist folder exists
# If it does, clear it
# If it doesn't, create it
if [ -d "dist" ]; then
	rm -rf dist
fi

mkdir dist

################################

echo "Building sizzle"
cd ./sizzle || exit
go build -o ../dist/sizzle
cd ../

echo "Building mkpath"
cd ./mkpath || exit
go build -o ../dist/mkpath
cd ../

echo "Building attic"
cd ./attic || exit
go build -o ../dist/attic
cd ../

echo "Building zn"
cd ./zn || exit
go build -o ../dist/zn
cd ../
