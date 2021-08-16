#!/bin/bash

#increment this before running the script if you plan to install the new version locally
VERSION="0.3.2"

ibmcloud plugin uninstall cra
sleep 1

if [ $? == '0' ]
then
    ibmcloud plugin install ./binaries/doi-cli-linux-amd64-0.3.2
fi
