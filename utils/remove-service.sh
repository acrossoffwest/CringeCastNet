#!/bin/bash

if [ -f /etc/systemd/system/cringecast.service ]
then
  systemctl stop cringecast.service
  systemctl disable cringecast.service
fi