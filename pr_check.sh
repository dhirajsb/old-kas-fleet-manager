#!/bin/bash -ex
#
# Copyright (c) 2019 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# This script is executed by a Jenkins job for each change request. If it
# doesn't succeed the change won't be merged.

# The version of `podman` available in the Jenkins environment doesn't work
# well when multiple sessions are created. The pause process that it creates is
# killed when the session finishes, but the file containing its pid is left
# around. That makes the next execution of `podman` fail because it can't join
# that pause process. To avoid that we need to explicitly tell `podman` and
# related tools to use a different directory to store state, and we also need
# to clean it before starting the build.
export XDG_RUNTIME_DIR="${PWD}/.xdg"
rm -rf "${XDG_RUNTIME_DIR}"
mkdir -p "${XDG_RUNTIME_DIR}"

# Set the `GOBIN` environment variable so that dependencies will be installed
# always in the same place, regardless of the value of `GOPATH`:
export GOBIN="${PWD}/.gobin"
export PATH="${GOBIN}:${PATH}"

export IMAGE_NAME="test/kas-fleet-manager"
export ENV MAS_SSO_CLIENT_ID="test-managed-services-api"
export MAS_SSO_CLIENT_SECRET="c1a79f76-272d-4b2b-ad27-2740fc81a508"
INTEGRATION_ENV="integration"

if [[ -z "${MAS_SSO_CLIENT_ID}" ]] || [[ -z "${MAS_SSO_CLIENT_SECRET}" ]];
then
   echo "Required mas sso env var: client id & client secret & crt is not provided"
   exit 1
fi

# copy dockerfile depending on targetted environment and set env vars in the dockerfile
if [[ -z "${OCM_ENV}" ]] || [[ "${OCM_ENV}" == "${INTEGRATION_ENV}" ]];
then
  cp docker/Dockerfile_template_mocked Dockerfile_integration_tests
else
  if [[ -z "${OCM_ENV}" ]] || [[ -z "${AWS_ACCESS_KEY}" ]] || [[ -z "${AWS_ACCOUNT_ID}" ]] || [[ -z "${AWS_SECRET_ACCESS_KEY}" ]] || [[ -z "${OCM_OFFLINE_TOKEN}" ]];  then
    echo "Required env var not provided. Exiting...".
    exit 1
  fi
  cp docker/Dockerfile_template Dockerfile_integration_tests
  sed -i "s/<ocm_env>/${OCM_ENV}/g" Dockerfile_integration_tests
  sed -i "s/<aws_access_key>/${AWS_ACCESS_KEY}/g" Dockerfile_integration_tests
  sed -i "s/<aws_account_id>/${AWS_ACCOUNT_ID}/g" Dockerfile_integration_tests
  sed -i "s/<aws_secret_access_key>/${AWS_SECRET_ACCESS_KEY}/g" Dockerfile_integration_tests
  sed -i "s/<ocm_offline_token>/${OCM_OFFLINE_TOKEN}/g" Dockerfile_integration_tests
fi

docker build -t "$IMAGE_NAME" -f Dockerfile_integration_tests .
docker run -i "$IMAGE_NAME"
