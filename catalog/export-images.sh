#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x

mkdir -p images

OS=$(uname -o)
if [ "${OS}" = "GNU/Linix" ]; then
    OS=Linux
fi
ARCH=$(uname -m)
if [ "${ARCH}" = "aarch64" ]; then
    ARCH=arm64
fi
curl -sL "https://github.com/google/go-containerregistry/releases/latest/download/go-containerregistry_${OS}_${ARCH}.tar.gz" >/tmp/go-containerregistry.tar.gz
tar -zxvf /tmp/go-containerregistry.tar.gz -C /tmp/
mv /tmp/crane images

CMD="./images/crane"

$CMD pull ghcr.io/appscode/kube-rbac-proxy:v0.11.0 images/appscode-kube-rbac-proxy-v0.11.0.tar
$CMD pull ghcr.io/appscode/kubectl:v1.24 images/appscode-kubectl-v1.24.tar
$CMD pull ghcr.io/kubestash/kubedump:v0.12.0 images/kubestash-kubedump-v0.12.0.tar
$CMD pull ghcr.io/kubestash/kubestash:v0.13.0 images/kubestash-kubestash-v0.13.0.tar
$CMD pull ghcr.io/kubestash/manifest:v0.1.0 images/kubestash-manifest-v0.1.0.tar
$CMD pull ghcr.io/kubestash/pvc:v0.12.0 images/kubestash-pvc-v0.12.0.tar
$CMD pull ghcr.io/kubestash/volume-snapshotter:v0.12.0 images/kubestash-volume-snapshotter-v0.12.0.tar
$CMD pull ghcr.io/kubestash/workload:v0.12.0 images/kubestash-workload-v0.12.0.tar

tar -czvf images.tar.gz images
