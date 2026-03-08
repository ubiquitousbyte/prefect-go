#!/bin/bash
set -e

# Script to fetch OpenAPI spec from a Prefect instance
# Usage: ./scripts/fetch-openapi.sh <version> [api-url]

VERSION=${1:-"latest"}
API_URL=${2:-"http://localhost:4200"}

if [ "$VERSION" = "latest" ]; then
    echo "Error: Please specify a version number (e.g., 3.6.21)"
    exit 1
fi

echo "Fetching OpenAPI spec from $API_URL for version $VERSION..."

# Create directory
mkdir -p "specs/$VERSION"

# Fetch the spec and format with jq
# First, try /api/openapi.json
if curl -f -s "$API_URL/api/openapi.json" | jq '.' > "specs/$VERSION/openapi.json" 2>/dev/null; then
    echo "✓ Fetched from $API_URL/api/openapi.json"
elif curl -f -s "$API_URL/openapi.json" | jq '.' > "specs/$VERSION/openapi.json" 2>/dev/null; then
    echo "✓ Fetched from $API_URL/openapi.json"
else
    echo "Error: Could not fetch OpenAPI spec from $API_URL"
    echo "Make sure:"
    echo "  1. Prefect server is running"
    echo "  2. The URL is correct"
    echo "  3. You can access the /docs endpoint"
    echo "  4. jq is installed (brew install jq)"
    exit 1
fi

# Verify the file was created and has content
if [ ! -s "specs/$VERSION/openapi.json" ]; then
    echo "Error: Downloaded file is empty or invalid"
    rm -f "specs/$VERSION/openapi.json"
    exit 1
fi

# Update symlink
rm -f specs/latest
ln -s "$VERSION" specs/latest

echo "✓ Spec saved to specs/$VERSION/openapi.json"
echo "✓ Updated specs/latest symlink"
echo ""
echo "Next steps:"
echo "  1. Review the spec: cat specs/$VERSION/openapi.json | jq '.info'"
echo "  2. Generate client: make generate"
