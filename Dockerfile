# base image
FROM resin/rpi-raspbian:latest

# copy files from host to container
COPY ./chatbot /

# main cmd start with docker run
ENTRYPOINT ["/chatbot"]
