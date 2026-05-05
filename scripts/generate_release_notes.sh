#!/usr/bin/env bash

set -euo pipefail

ROOT=$(git rev-parse --show-toplevel)
OUTPUT="$ROOT/extensions/release-notes/release_notes.md"
TMP="$ROOT/extensions/release-notes/tmp_release_notes.md"

: >"$OUTPUT"

TAGS=$(git tag --sort=creatordate)
# gets all existing tags for deterministic content
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

	cat "$TMP" "$OUTPUT" >"${OUTPUT}.new" && mv "${OUTPUT}.new" "$OUTPUT"
done

# creates release notes for the non-existing upcoming version
# and appends it at the top of the file
if [[ $# -ge 1 ]]; then
	NEW_TAG="$1"
	echo "Preparing release notes for upcoming tag: $NEW_TAG..."

	NOTES=$(gh api repos/opentelekomcloud/provider-opentelekomcloud/releases/generate-notes \
		-f tag_name="$NEW_TAG" \
		--jq .body)

	{
		echo "# $NEW_TAG"
		echo ""
		echo "$NOTES"
		echo ""
	} >"$TMP"

	cat "$TMP" "$OUTPUT" >"${OUTPUT}.new" && mv "${OUTPUT}.new" "$OUTPUT"
fi

rm -f "$TMP"

echo "Done. Output in $OUTPUT"
