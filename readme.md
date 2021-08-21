# Go docker-compose

Go Docker image linked to a docker-compose.yml file with hot reload provided by [air](https://github.com/cosmtrek/air).

## Starting the container

```bash
docker-compose up
```

## Writing code

All the code inside `backend/` folder is mounted inside the container and monitored by Air. Every time a change is detected, Air will recompile and execute the code.
