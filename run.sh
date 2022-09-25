#!/bin/sh

docker build -t twitter . && docker run -p 9191:8080 twitter