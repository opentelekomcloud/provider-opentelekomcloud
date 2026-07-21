#!/usr/bin/env bash
#
# Nightly test runner for UPTEST examples
# Reads configuration from nightly_test_config.yaml and runs tests
# for each configured namespace/service/test_case.yaml with optional custom settings.
#
# Usage:
#   ./scripts/nightly_test.sh [config_path]
#
# Environment variables:
#   UPTEST_CLOUD_CREDENTIALS  - Cloud credentials for testing (required)
#   UPTEST_DATASOURCE_PATH    - Path to datasource YAML file
#   UPTEST_TEST_DIR           - Test directory for controlplane dumps
#   UPTEST_UPDATE_PARAMETER   - Update parameter (default: "")
#   UPTEST_DEFAULT_TIMEOUT    - Global default timeout
#   UPTEST_SKIP_IMPORT        - Skip import testing (default from config)
#   UPTEST_NIGHTLY_CONFIG     - YAML config content (alternative to file argument)

set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Determine config path from argument or env var
CONFIG_PATH="${1:-}"
CONFIG_CONTENT="${UPTEST_NIGHTLY_CONFIG:-}"
TEMP_CONFIG_FILE=""

if [[ -n "$CONFIG_PATH" && -f "$CONFIG_PATH" ]]; then
	# Use provided file path
	:
elif [[ -n "$CONFIG_CONTENT" ]]; then
	# Create temp file in repo
	TEMP_CONFIG_FILE="$REPO_ROOT/.work/nightly_config_temp.yaml"
	mkdir -p "$(dirname "$TEMP_CONFIG_FILE")"
	echo "$CONFIG_CONTENT" >"$TEMP_CONFIG_FILE"
	CONFIG_PATH="$TEMP_CONFIG_FILE"
	echo ">>> Created temporary config file: $CONFIG_PATH"
else
	echo "ERROR: No config file or UPTEST_NIGHTLY_CONFIG env var provided."
	exit 1
fi

# Verify config file exists
if [[ ! -f "$CONFIG_PATH" ]]; then
	echo "ERROR: Config file not found: $CONFIG_PATH"
	exit 1
fi

# Verify yq is available
if ! command -v yq &>/dev/null; then
	echo "ERROR: yq is required but not installed. Please install yq first."
	exit 1
fi

cd "$REPO_ROOT" || exit 1

REPORT_DIR=".work/nightly-report"
mkdir -p "$REPORT_DIR"

REPORT_FILE="$REPORT_DIR/report.txt"
TIMESTAMP=$(date -u '+%Y-%m-%d %H:%M:%S UTC')

echo "============================================"
echo " Nightly UPTEST Runner"
echo "============================================"
echo "Config: $CONFIG_PATH"
echo "Repo:   $REPO_ROOT"
echo "Date:   $TIMESTAMP"
echo "============================================"
echo ""

# Print the full config
echo ">>> Full configuration:"
echo "--------------------------------------------"
cat "$CONFIG_PATH"
echo "--------------------------------------------"
echo ""

# Initialize report
cat >"$REPORT_FILE" <<EOF
============================================
 Nightly UPTEST Test Report
============================================
Date: $TIMESTAMP
Config: $CONFIG_PATH
============================================

EOF

# Build and deploy the provider once before all tests
echo ">>> Building and deploying the provider..."
make local-deploy
echo ""

# Get the number of namespaces
NS_COUNT=$(yq eval '.namespaced | length' "$CONFIG_PATH")
echo ">>> Found $NS_COUNT namespace(s) to test"
echo ""

FAILED_NAMESPACES=()
PASSED_NAMESPACES=()

for i in $(seq 0 $((NS_COUNT - 1))); do
	NAMESPACE=$(yq eval ".namespaced[$i] | keys[0]" "$CONFIG_PATH")

	# Check if namespace is enabled
	ENABLED_RAW=$(yq eval ".namespaced[$i].${NAMESPACE}.enabled" "$CONFIG_PATH")
	if [[ "$ENABLED_RAW" == "null" ]]; then
		ENABLED="true"
	else
		ENABLED="$ENABLED_RAW"
	fi

	if [[ "$ENABLED" == "false" ]]; then
		echo ">>> Skipping disabled namespace: $NAMESPACE"
		continue
	fi

	TEST_COUNT=$(yq eval ".namespaced[$i].${NAMESPACE}.tests | length" "$CONFIG_PATH")
	IMPORT_TEST_RAW=$(yq eval ".namespaced[$i].${NAMESPACE}.import_test" "$CONFIG_PATH")
	if [[ "$IMPORT_TEST_RAW" == "null" ]]; then
		IMPORT_TEST="true"
	else
		IMPORT_TEST="$IMPORT_TEST_RAW"
	fi
	TIMEOUT=$(yq eval ".namespaced[$i].${NAMESPACE}.timeout // \"20m\"" "$CONFIG_PATH")

	echo "============================================"
	echo " Testing namespace: $NAMESPACE ($TEST_COUNT test file(s))"
	echo " Import test: $IMPORT_TEST"
	echo " Timeout: $TIMEOUT"
	echo "============================================"

	echo "Namespace: $NAMESPACE" >>"$REPORT_FILE"
	echo "Tests: $TEST_COUNT | Import: $IMPORT_TEST | Timeout: $TIMEOUT" >>"$REPORT_FILE"
	echo "============================================" >>"$REPORT_FILE"
	echo "" >>"$REPORT_FILE"

	NAMESPACE_FAILED=false

	for j in $(seq 0 $((TEST_COUNT - 1))); do
		TEST_FILE=$(yq eval ".namespaced[$i].${NAMESPACE}.tests[$j]" "$CONFIG_PATH")
		EXAMPLE_PATH="examples/namespaced/${NAMESPACE}/${TEST_FILE}"

		echo ""
		echo ">>> Running test: $EXAMPLE_PATH"
		echo "    Import test: $IMPORT_TEST | Timeout: $TIMEOUT"

		echo "  Test: $EXAMPLE_PATH" >>"$REPORT_FILE"

		export UPTEST_EXAMPLE_LIST="$EXAMPLE_PATH"
		export UPTEST_DEFAULT_TIMEOUT="$TIMEOUT"

		if [[ "$IMPORT_TEST" == "false" ]]; then
			export UPTEST_SKIP_IMPORT="true"
		else
			unset UPTEST_SKIP_IMPORT
		fi

		# Run the test
		if make uptest 2>&1 | tee "/tmp/uptest-${NAMESPACE}-${TEST_FILE}.log"; then
			echo ">>> PASSED: $EXAMPLE_PATH"
			echo "    Status: PASSED" >>"$REPORT_FILE"
		else
			echo ">>> FAILED: $EXAMPLE_PATH"
			NAMESPACE_FAILED=true
			echo "    Status: FAILED" >>"$REPORT_FILE"
		fi
		echo "" >>"$REPORT_FILE"
	done

	if [[ "$NAMESPACE_FAILED" == "true" ]]; then
		FAILED_NAMESPACES+=("$NAMESPACE")
		echo ""
		echo "!!! Namespace $NAMESPACE had failures"
		echo "  Result: FAILED" >>"$REPORT_FILE"
	else
		PASSED_NAMESPACES+=("$NAMESPACE")
		echo ""
		echo ">>> Namespace $NAMESPACE passed all tests"
		echo "  Result: PASSED" >>"$REPORT_FILE"
	fi
	echo "" >>"$REPORT_FILE"
done

# Summary
echo "============================================"
echo " Nightly Test Summary"
echo "============================================"
echo "Passed: ${#PASSED_NAMESPACES[@]} namespace(s)"
for ns in "${PASSED_NAMESPACES[@]}"; do
	echo "  - $ns"
done
echo ""
echo "Failed: ${#FAILED_NAMESPACES[@]} namespace(s)"
for ns in "${FAILED_NAMESPACES[@]}"; do
	echo "  - $ns"
done
echo ""
echo "============================================"

# Write summary to report
cat >>"$REPORT_FILE" <<EOF

============================================
 Nightly Test Summary
============================================
Passed: ${#PASSED_NAMESPACES[@]} namespace(s)
EOF

for ns in "${PASSED_NAMESPACES[@]}"; do
	echo "  - $ns" >>"$REPORT_FILE"
done

echo "" >>"$REPORT_FILE"
echo "Failed: ${#FAILED_NAMESPACES[@]} namespace(s)" >>"$REPORT_FILE"

for ns in "${FAILED_NAMESPACES[@]}"; do
	echo "  - $ns" >>"$REPORT_FILE"
done

echo "" >>"$REPORT_FILE"
echo "============================================" >>"$REPORT_FILE"

if [[ ${#FAILED_NAMESPACES[@]} -gt 0 ]]; then
	echo "RESULT: Some tests FAILED" >>"$REPORT_FILE"
	echo ""
	echo "RESULT: Some tests FAILED"
	exit 1
else
	echo "RESULT: All tests PASSED" >>"$REPORT_FILE"
	echo ""
	echo "RESULT: All tests PASSED"
	exit 0
fi
