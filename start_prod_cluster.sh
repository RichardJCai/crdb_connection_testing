#!/bin/bash -i

ARG2=${2:-richardc-test}
ARG3=${3:-v21.1.0-alpha.3}

shopt -s expand_aliases

export CLUSTER="${ARG2}"
roachprod create ${CLUSTER} -n $1 --local-ssd

ssh-add ~/.ssh/google_compute_engine

roachprod stage ${CLUSTER} workload
roachprod stage ${CLUSTER} release ${ARG3}

roachprod start ${CLUSTER}
