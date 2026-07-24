#!/usr/bin/env bash
# should run from repo root
set -euo pipefail

CRD_DIR="package/crds"

if ! git rev-parse --is-inside-work-tree &>/dev/null; then
	echo "Error: not a git repository"
	exit 1
fi

changed_files=$(git status --porcelain -- "$CRD_DIR" || true)

if [[ -z "$changed_files" ]]; then
	echo "No changes detected in $CRD_DIR"
	echo "Are you in repos root dir?"
	exit 0
fi

echo "Changes detected in $CRD_DIR:"
echo "$changed_files"
echo ""

while IFS= read -r line; do
	file=$(echo "$line" | awk '{print $2}')
	full_path="$file"

	if [[ $(echo "$line" | awk '{print $1}') == "R*" ]]; then
		full_path=$(echo "$line" | awk '{print $3}')
	fi

	if [[ -f "$full_path" ]]; then
		echo "Applying: $full_path"
		kubectl apply -f "$full_path"
	else
		echo "Skipping (file not found): $full_path"
	fi
done <<<"$changed_files"

echo ""
echo "Done applying CRDs"
