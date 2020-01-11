#!/bin/sh
curl https://gp.ynov-bordeaux.com/assets/superhero/all.json | jq '.[] | select(.id==70) | .name'
