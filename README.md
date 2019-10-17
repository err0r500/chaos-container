# Chaos-container

This simple container can be used in your chaos engineering experiments

## Status code

- get returns a 500 on every request

```
docker run -p 8080:8080 err0r500/chaos-container
```

- status env var can be used in order to customize the response code 

```
docker run -p 8080:8080 -e STATUS=503 err0r500/chaos-container
```
