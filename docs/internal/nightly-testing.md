# Nightly E2E Testing

The nightly tests run all `examples/namespaced/*.yaml` UPTEST cases.

## Configuration

The test config lives in GitHub Actions variable `UPTEST_NIGHTLY_CONFIG` (YAML). It defines which namespaces to test, per-namespace `enabled`/`import_test`/`timeout` settings.

## GitHub Actions Workflow

`.github/workflows/e2e-nightly.yaml` triggers manually or daily at midnight UTC. It validates the config variable, builds and deploys the provider, then runs `scripts/nightly_test.sh`. Test reports and cluster dumps are uploaded as artifacts.

## Script

The script reads the YAML config using `yq`, iterates over each namespace, and skips any with `enabled: false`. For each test file it sets `UPTEST_EXAMPLE_LIST`, `UPTEST_DEFAULT_TIMEOUT`, and optionally `UPTEST_SKIP_IMPORT` env vars, then calls `make uptest`. All tests run even if some fail, and a summary report is written to `.work/nightly-report/report.txt`.

