#!/usr/bin/bash

# source the dotenv file.
set -a            
source .env
set +a

# set the values within local.yaml
envsubst < specs/template.y*ml > deploy.yaml 