#!/bin/bash

if [ -f /etc/systemd/system/cringecast.service ]
then
    if (systemctl -q is-active cringecast.service)
    then
      systemctl stop cringecast.service
    fi
    systemctl disable cringecast.service
    rm /etc/systemd/system/cringecast.service
fi