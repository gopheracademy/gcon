#!/bin/bash

set -eo

go build
docker login -e="$DOCKER_EMAIL" -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"
docker build -t $IMAGE:$COMMIT .
docker tag $IMAGE:$COMMIT $IMAGE:latest
docker tag $IMAGE:$COMMIT $IMAGE:travis-$TRAVIS_BUILD_NUMBER
docker push $IMAGE
curl https://sdk.cloud.google.com | bash
source /home/travis/.bashrc
gcloud components update kubectl
kubectl get pods
kubectl patch deployment $DEPLOYMENT -p '{"spec":{"template":{"spec":{"containers":[{"name":"'"$CONTAINER_NAME"'","image":"'"$IMAGE_VERSION"'"}]}}}}'
