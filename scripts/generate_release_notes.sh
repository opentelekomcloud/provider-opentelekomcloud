#!/usr/bin/env bash

set -euo pipefail
ROOT=$(git rev-parse --show-toplevel)
OUTPUT="$ROOT/extensions/release-notes/release_notes.md"
TMP="$ROOT/extensions/release-notes/tmp_release_notes.md"

: >"$OUTPUT"

TAGS=$(git tag --sort=creatordate)

for TAG in $TAGS; do
	echo "Processing $TAG..."

	NOTES=$(gh api repos/opentelekomcloud/provider-opentelekomcloud/releases/generate-notes \
		-f tag_name="$TAG" \
		--jq .body)

	{
		echo "# $TAG"
		echo ""
		echo "$NOTES"
		echo ""
	} >"$TMP"

	cat "$TMP" "$OUTPUT" >"${OUTPUT}.new" &&
		mv "${OUTPUT}.new" "$OUTPUT"
done

rm -f "$TMP"

echo "Done. Output in $OUTPUT"
