#!/bin/bash

docker rm mongodb
docker run --name mongo -d -p 27017:27017 mongo
