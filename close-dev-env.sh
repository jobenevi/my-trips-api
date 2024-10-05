#!/bin/bash
sudo docker-compose -f docker-compose-dev.yml down -v
sudo docker volume prune -f
sudo docker network prune -f
sudo docker image prune -f

sudo rm -rf $HOME/data/postgres-data

sudo docker container prune -f
sudo docker image prune -a -f
sudo docker volume prune -f
sudo docker network prune -f

sudo fuser -k 5432/tcp