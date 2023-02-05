#!/bin/bash
sudo docker pull a48zhang/melting:dev
sudo docker rm "$(sudo docker stop "$(sudo docker ps -f ancestor=a48zhang/melting:dev -q)")"
sudo docker run -dp 65000:65000 --env-file ./melting.env --name=melting_dev a48zhang/melting:dev
