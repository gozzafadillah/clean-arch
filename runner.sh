#!/bin/bash

##build
docker build --tag $1:$2 .

##run
docker run --rm -d --name runner-$1-$2 -p 19000:8080 --env-file ./config.env $1:$2

##running - compose
## docker-compose up --build