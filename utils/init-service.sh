#!/bin/bash

if [ ! -f ./cringecast.service ]
then
    cp ./cringecast.service.example ./cringecast.service
    echo "Configure service file: ./cringecast.service, and run this script again"
    return
else
  if [ -f /etc/systemd/system/cringecast.service ]
  then
      rm /etc/systemd/system/cringecast.service
  fi
  ln -s $(pwd)/cringecast.service /etc/systemd/system/cringecast.service
  systemctl enable cringecast.service
  systemctl start cringecast.service
fi