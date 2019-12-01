# Vault-Abstraction

## Overview
Abstraction for vault smart contracts


### Running the service
```shell
make serve
```

### Running Unit Tests
```shell
make test
```

### Linting
```shell
make lint
```

### Docker build
```shell
make docker-build
```

### User Testing
```shell
curl 0.0.0.0:8080/api/v1/health/status
```

### Swagger UI
Swagger ui will be exposed at:
```
0.0.0.0:8080/api/v1/swagger/docs/swaggerui/
```

### Endpoints

```
curl -H "Content-Type: application/json" -X POST -d '{"ipfsHash":"QmW3FgNGeD46kHEryFUw1ftEUqRw254WkKxYeKaouz7DJA", "name":"hope"}' 0.0.0.0:8080/ipfs-media-platform/manager/add-vault
```

curl -H "accept: application/json" -H "Content-Type: application/json" -X POST -d 'ipfsHash=123dsad&name=name' 0.0.0.0:8080/ipfs-media-platform/manager/add-vault