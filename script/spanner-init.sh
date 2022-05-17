#!/usr/bin/env bash
set -eu

sleep 10
gcloud config configurations create emulator
gcloud config set auth/disable_credentials true
gcloud config set project ${SPANNER_PROJECT_ID}
gcloud config set api_endpoint_overrides/spanner ${SPANNER_EMULATOR_URL}
gcloud spanner instances create ${SPANNER_INSTANCE_ID} --config=emulator-config --description=Emulator --nodes=1
gcloud spanner databases create ${SPANNER_DATABASE_ID} --instance=${SPANNER_INSTANCE_ID}