#!/usr/bin/env bash
# cleanup.sh
# Description: Cleanup buils files
# Author: BlueFalconHD

# If the dist directory exists, remove it
# If the dist directory does not exist, do nothing
if [ -d "dist" ]; then
	rm -r dist
fi

# Tidy project files
cd ./mkpath || exit
go mod tidy
cd ../

cd ./sizzle/ || exit
go mod tidy
cd ../

cd ./attic || exit
go mod tidy
cd ../

cd ./zn || exit
go mod tidy
cd ../
