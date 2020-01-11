#!/bin/sh

curl -s https://gp.ynov-bordeaux.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"JROUZIERE\"}}){id}}"}' | cut -c24-25
