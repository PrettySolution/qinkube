1. `kubectl get --raw /openapi/v2  > k8s-openapi-v2.json`
2. ```shell
docker run \
    -v ./k8s-openapi-v2.json:/app/swagger.json \
    -p 8081:8080 \
    swaggerapi/swagger-ui
```