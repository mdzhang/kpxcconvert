#!/bin/bash
set -eu -o -x pipefail

VERSION="$(git describe --tags --abbrev=0)"
BRANCH="$(git rev-parse --abbrev-ref HEAD)"

LAST_VERSION="$(git describe --tags $(git rev-list --tags --max-count=1))"

DESCRIPTION="# Bump ${LAST_VERSION} → ${VERSION}"

if [ "${VERSION}" != "${LAST_VERSION}" ]; then
  DESCRIPTION="${DESCRIPTION}<br/>**Changelog**<br/>$(git log ${LAST_VERSION}..HEAD --oneline --no-merges)"
else
  DESCRIPTION="# Bump N/A → ${VERSION}<br/>**Changelog**<br/>$(git log ${VERSION} --oneline --no-merges)"
fi

github-release mdzhang/kpxcconvert "${VERSION}" "${BRANCH}" "${DESCRIPTION}" "dist/kpxcconvert-${VERSION}-*"
