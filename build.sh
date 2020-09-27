#!/usr/bin/env bash
docker build -t sillyhatxu/web-service-demo -f application-api/Dockerfile .
docker tag sillyhatxu/web-service-demo:latest sillyhatxu/web-service-demo:api-1.0
docker push sillyhatxu/web-service-demo:api-1.0

docker build -t sillyhatxu/web-service-demo -f application-internal-api/Dockerfile .
docker tag sillyhatxu/web-service-demo:latest sillyhatxu/web-service-demo:internal-api-1.0
docker push sillyhatxu/web-service-demo:internal-api-1.0