#!/bin/bash

#install zip package and add it to PATH
#run with admin privileges

# loop through all lambda directories
for d in *-lambda ; do
  cd "$d" #cd into directory
  # build lambda
  GOOS=linux GOARCH=amd64 go build -o ../lambda_build/main main.go
  # cd into folder and zip file
  cd ../lambda_build
  zip "$d"1.0.0.zip ./main
  echo lambda "$d" built

  cd ..
done