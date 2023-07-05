#!/bin/bash

##Docker objects cleaner
# docker stop $(docker ps -qa) && docker system prune -af --volumes
##Adjust permission to posgres-data (host volume) and run compose
sudo chmod -R 777 ./postgres-data
docker compose up