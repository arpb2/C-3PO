#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# This check verifies consistency across the structure of the backend. Disallowing any dependency between layers that
# shouldn't know each other (eg. Presentation using data sources explicitly)

# Check that the data files don't depend on presentation ones
diff=$(grep -iRl "github.com/arpb2/C-3PO/pkg/presentation" ./pkg/data || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Data files shouldn't depend on presentation ones" >&2
  exit 1
fi

# Check that the presentation files don't depend on the data ones
diff=$(grep -iRl "github.com/arpb2/C-3PO/pkg/data" ./pkg/presentation || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Presentation files shouldn't depend on data ones" >&2
  exit 1
fi

# Check that the domain files don't depend on the presentation/data ones
diff=$(grep -iERl "github.com/arpb2/C-3PO/pkg/(presentation|data)" ./pkg/domain || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Domain files shouldn't depend on data / presentation ones" >&2
  exit 1
fi

# Check that the pkg files don't depend on cmd ones
diff=$(grep -iRl "github.com/arpb2/C-3PO/cmd" ./pkg/ || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Pkg files shouldn't depend on cmd ones" >&2
  exit 1
fi

# Check that the test (mock) files don't depend on presentation / data ones
diff=$(grep -iERl "github.com/arpb2/C-3PO/pkg/(presentation|data)" ./test || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Test mock files shouldn't depend on data / presentation ones" >&2
  exit 1
fi

# Check that the test (mock) files don't depend on cmd ones
diff=$(grep -iERl "github.com/arpb2/C-3PO/cmd" ./test || true)
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Test mock files shouldn't depend on cmd ones" >&2
  exit 1
fi