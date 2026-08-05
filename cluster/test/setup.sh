#!/usr/bin/env bash
set -aeuo pipefail

MARKER=/tmp/cluster-provider-setup.done

echo "Running setup.sh"

if [[ -f $MARKER ]]; then
	echo "Already ran, skipping to save time..."
	exit 0
else

	echo "Creating cloud credential secrets..."

	### update default MRAP to only activate modern APIs

	cat <<EOF | ${KUBECTL} apply -f -
apiVersion: apiextensions.crossplane.io/v1alpha1
kind: ManagedResourceActivationPolicy
metadata:
  name: default
spec:
  activate:
  - "*.opentelekomcloud.m.crossplane.io"
EOF

	### test examples should use `test` namespace
	cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: test
EOF

	${KUBECTL} -n test create secret generic provider-secret-de \
		--from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS_DE}" \
		--dry-run=client -o yaml | ${KUBECTL} apply -f -

	# needed for testing cross-region-replication
	${KUBECTL} -n test create secret generic provider-secret-nl \
		--from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS_NL}" \
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
      name: provider-secret-de
      namespace: test
      key: credentials
EOF

	echo "Creating a provider config for second region..."
	cat <<EOF | ${KUBECTL} apply -f -
apiVersion: opentelekomcloud.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default-nl
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret-nl
      namespace: test
      key: credentials
EOF

	echo "Kind clusters need some time to process new CRDs..."
	echo "Give them a few sec to be ready..."
	sleep 30s

	${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
	${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

	touch "$MARKER"
fi
