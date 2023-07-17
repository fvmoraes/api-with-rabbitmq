#!/bin/bash

##Docker objects cleaner
# docker stop $(docker ps -qa) && docker system prune -af --volumes
##Adjust permission to posgres_data and rabbitmq_data (host volume) and run compose
sudo chmod -R 777 ./build/volumes/*
sudo chmod 400 ./build/volumes/rabbitmq_data/.erlang.cookie
docker compose up