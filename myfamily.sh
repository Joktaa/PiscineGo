#!/bin/sh

curl -s https://gp.ynov-bordeaux.com/assets/superhero/all.json | jq ".[] | select(.id==$HERO_ID) | .connections.relatives" | tr -d '"'
