#!/bin/bash

openssl rand -base64 756 > env/dev/dev-keyfile
chmod 400 env/dev/dev-keyfile
chown 999:999 env/dev/dev-keyfile
