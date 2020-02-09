#! /bin/bash

echo "------------build the nginx docker images------------"
docker build -t wts_nginx ./nginx
echo "------------finished------------"

echo "------------build the server docker images------------"
docker build -t wts_server ./server
echo "------------finished------------"

echo "------------pull the mysql image------------"
docker pull mysql:8.0.18
echo "------------finished------------"


echo "------------export and zip the images to deploy------------"
docker save -o ./wts_server.tar wts_nginx:latest wts_server:latest mysql:8.0.18
gzip wts_server.tar
cp -f wts_server.tar.gz ./deploy
rm wts_server.tar.gz
echo "------------finished------------"

echo "------------clean up the docker build environment------------"
docker image rm wts_nginx:latest wts_server:latest ubuntu:18.04 mysql:8.0.18
echo "------------finished------------"
