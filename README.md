# Docker-GoLang1.6-Server

# Start your local docker VM
* docker-machine start default

# Build the container
* docker build -t test-api .

# Run the container
* docker run --publish 3010:8080 --name test-api --rm test-api

# Visit your new api
* open "$(docker-machine ip default):3010"
