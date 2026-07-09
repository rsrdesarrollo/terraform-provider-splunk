#!/bin/bash

set -e

BUMP_TYPE=${1:-patch}

# Get the latest tag
LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
LATEST_VERSION=${LATEST_TAG#v}

# Parse version
IFS='.' read -r MAJOR MINOR PATCH <<< "$LATEST_VERSION"

echo "Current version: v$MAJOR.$MINOR.$PATCH"

case $BUMP_TYPE in
    patch)
        NEW_PATCH=$((PATCH + 1))
        NEW_VERSION="$MAJOR.$MINOR.$NEW_PATCH"
    ;;
    minor)
        NEW_MINOR=$((MINOR + 1))
        NEW_VERSION="$MAJOR.$NEW_MINOR.0"
    ;;
    major)
        NEW_MAJOR=$((MAJOR + 1))
        NEW_VERSION="$NEW_MAJOR.0.0"
    ;;
    *)
        echo "Invalid bump type: $BUMP_TYPE. Use: patch, minor, or major"
        exit 1
    ;;
esac

NEW_TAG="v$NEW_VERSION"

echo "New version: $NEW_TAG"

# Update CHANGELOG.md
CHANGELOG_ENTRY="## $NEW_VERSION"
sed -i '' "1s/^/$CHANGELOG_ENTRY\n* \n\n/" CHANGELOG.md

echo "Updated CHANGELOG.md"

git add CHANGELOG.md
git commit -m "bump version to $NEW_TAG, update changelog"

# Create tag
git tag -a "$NEW_TAG" -m "Release $NEW_VERSION"
echo "Created tag: $NEW_TAG"

# Push tag
git push origin master
git push origin "$NEW_TAG"
echo "Pushed tag: $NEW_TAG"

echo ""
echo "✅ Version bumped successfully to $NEW_VERSION"
