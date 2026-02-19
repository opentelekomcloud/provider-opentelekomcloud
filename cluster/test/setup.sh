#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."

### test examples should use `test` namespace
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: test
EOF

${KUBECTL} -n test create secret generic provider-secret \
	--from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" \
	--dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: test
      key: credentials
EOF

echo "Kind clusters need some time to process new CRDs..."
echo "Give them a minute to be ready..."
sleep 1m

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
