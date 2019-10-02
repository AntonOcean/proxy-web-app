#!/usr/bin/env bash

openssl req \
    -newkey rsa:2048 \
    -x509 \
    -nodes \
    -keyout proxy-web-app.key \
    -new \
    -out proxy-web-app.crt \
    -subj /CN=localhost \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
    -sha256 \
    -days 3650
