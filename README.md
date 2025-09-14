## Crowd Monitoring Broker

## Note

Before running the `docker` services you must have a domain that has valid `ssl` certificates and configures the `docker-compose.yml` file and change the `hostname` of the cetificates.

## Guide 
1. Clone the repository using `git` or `download` manually
2. Compile the source into executable using `go build -o <your_executable_name> main.go`
3. Run the executable using `./your_executable`

## Run Dockerized MYSQL and MQTT 
1. Go to the repo root directory
2. Run docker command `docker-compose up` and `docker-compose up -d`
3. Verify docker container is running `docker-compose ps`

## Any Bugs?
If your encounter any issues with the script please submit a ticket.
