#!/bin/bash

# binaries
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s"
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s"

# dirs
rm -rf ./build
mkdir ./build
mkdir ./build/_db
mkdir ./build/public

# add execution perms
chmod +x gasonline

#copy files and dirs
cp appconfig.toml ./build
mv gasonline.exe ./build
mv gasonline ./build
cp -R wui/dist/* ./build/public