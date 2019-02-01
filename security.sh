#!/bin/bash

gosec -fmt=json -out=gosec.json ./...
cat gosec.json
