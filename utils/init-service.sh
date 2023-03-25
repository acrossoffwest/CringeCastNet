#!/bin/bash

if [ ! -f ./cringecast.service ]
then
    cp ./cringecast.service.example ./cringecast.service
    echo "Configure service file: ./cringecast.service, and run this script again"
    return
else
  ln -s $(pwd)/cringecast.service /etc/systemd/system/cringecast.service
  sudo systemctl enable cringecast.service
  sudo systemctl start cringecast.service
fi