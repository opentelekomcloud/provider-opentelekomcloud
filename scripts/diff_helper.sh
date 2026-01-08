#!/usr/bin/env bash
set -euo pipefail

ROOT="$(git rev-parse --show-toplevel 2>/dev/null)"

if [[ -z "$ROOT" ]]; then
	echo "not inside a git repository" >&2
	exit 1
fi

BASE="${1:-origin/HEAD}"

cd "$ROOT"

git diff "$BASE" -- \
	':(exclude)**/zz_*' \
	':(exclude)package/crds/**' \
	':(exclude)examples-generated/**'
