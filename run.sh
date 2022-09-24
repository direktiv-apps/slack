#!/bin/sh

docker build -t slack . && docker run -p 9191:8080 slack