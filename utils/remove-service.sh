#!/bin/bash

if [ -f /etc/systemd/system/cringecast.service ]
then
    systemctl disable cringecast.service
fi