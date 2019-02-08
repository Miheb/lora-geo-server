# Start dev container (run once)
docker run -d -it --volume=$PWD:/go/src/github.com/brocaar/lora-geo-server --name lora_geo_dev khbx/lora_geo_dev:latest

# Requirements (run once)
sudo docker exec lora_geo_dev /bin/sh -c "make dev-requirements && make requirements"

# Compile and recreate imag
sudo docker exec lora_geo_dev /bin/sh -c "make clean && make build" && sudo docker build -t khbx/lora_geo -f Dockerfile-khbx .

# Image was built using
docker build -f Dockerfile-devel -t khbx/lora_geo_dev .
