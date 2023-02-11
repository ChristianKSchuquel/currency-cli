#!/bin/bash

go build main.go

mv main currency

cp currency /usr/local/bin
