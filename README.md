# Description
Simple Go lang app which provides a rest api layer to perform CRUD operations on banking meta data.

# Docker flow
```
docker-compose build && docker-compose up -d --remove-orphans
```

# Run locally
```
make run
```

# Endpoints
http://127.0.0.1:8080/account - GET, POST, DELETE
http://127.0.0.1:8080/account/[id] - GET