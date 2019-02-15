#!/bin/sh

docker kill $(docker ps -q)
docker pull nyfanilo/monitor360:latest
docker run -e "ALLOWED_HOST=assistant.rafanilo.xyz" -e "ENV=production" -p 443:443 -p 80:80 -d nyfanilo/monitor360:latest