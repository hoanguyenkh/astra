# syntax=docker/dockerfile:1

## Build Image
FROM golang:1.18-bullseye as build

WORKDIR /astra
COPY . /astra

RUN BUILD_TAGS=muslc make build-e2e-chain-init

## Deploy image
FROM ubuntu

COPY --from=build /astra/build/chain_init /bin/chain_init

ENV HOME /astra
WORKDIR $HOME

ENTRYPOINT [ "chain_init" ]
