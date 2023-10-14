#!/bin/sh
echo "Stage 1: build user"
cd user 
docker build . -t user:1.0 
cd ..

echo "Stage 2: build admin"
cd admin
docker build . -t admin:1.0
cd ..

echo "Stage 3: deploy"
cd deployments
docker compose up 