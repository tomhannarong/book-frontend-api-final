start dev

-go mod tidy
-go run ./server.go

start docker 

- docker-compose build

start step by step

2. // start 
-docker-compose up -d



start swarm on production step by step

1. //  comment command
placement                   | in docker-compose.yml
max_replicas_per_node: 1    | in docker-compose.yml

2. // start swarm
- docker stack deploy -c docker-compose.yml book

2. // stop swarm
- docker service ls
- docker service rm XX -> id service