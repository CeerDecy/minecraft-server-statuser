#!/bin/bash

TAG=$1

docker buildx build -t ceerdecy/minecraft-server-statuser:"${TAG}" .