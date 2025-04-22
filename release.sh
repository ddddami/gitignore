#!/bin/bash
APP_NAME="gitignore"
VERSION="v0.1.0"
OUTPUT_DIR="dist"
FILES_TO_INCLUDE=("README.md" "LICENSE")
CHECKSUM_FILE="${APP_NAME}_${VERSION}_checksums.txt"
PLATFORMS=(
  "linux/amd64"
  "linux/arm64"
  "darwin/amd64"
  "darwin/arm64"
  "windows/amd64"
  "windows/arm64"
)

mkdir -p $OUTPUT_DIR

echo "# ${APP_NAME} ${VERSION} checksums" > "$OUTPUT_DIR/$CHECKSUM_FILE"
echo "Generated on $(date)" >> "$OUTPUT_DIR/$CHECKSUM_FILE"
echo "" >> "$OUTPUT_DIR/$CHECKSUM_FILE"
echo "## SHA256" >> "$OUTPUT_DIR/$CHECKSUM_FILE"
echo '```' >> "$OUTPUT_DIR/$CHECKSUM_FILE"

for PLATFORM in "${PLATFORMS[@]}"
do
  IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
  OUTPUT_NAME="${APP_NAME}_${VERSION}_${GOOS}_${GOARCH}"
  BINARY_NAME="${APP_NAME}"
  ARCHIVE_EXT="tar.gz"
  
  if [ "$GOOS" = "windows" ]; then
    BINARY_NAME="${APP_NAME}.exe"
    ARCHIVE_EXT="zip"
  fi
  
  echo "Building for $GOOS/$GOARCH..."
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_DIR/$OUTPUT_NAME/$BINARY_NAME" ./cmd/gitignore
  
  cp "${FILES_TO_INCLUDE[@]}" "$OUTPUT_DIR/$OUTPUT_NAME/"
  
  if [ "$GOOS" = "windows" ]; then
    echo "Packaging $OUTPUT_NAME.zip..."
    (cd "$OUTPUT_DIR" && zip -r "$OUTPUT_NAME.zip" "$OUTPUT_NAME")
  else
    echo "Packaging $OUTPUT_NAME.tar.gz..."
    tar -czvf "$OUTPUT_DIR/$OUTPUT_NAME.tar.gz" -C "$OUTPUT_DIR" "$OUTPUT_NAME"
  fi
  
  # Generate SHA256 checksum
  ARTIFACT_FILE="$OUTPUT_DIR/$OUTPUT_NAME.$ARCHIVE_EXT"
  if [[ "$OSTYPE" == "darwin"* ]]; then
    SHA256=$(shasum -a 256 "$ARTIFACT_FILE" | awk '{print $1}')
  else
    SHA256=$(sha256sum "$ARTIFACT_FILE" | awk '{print $1}')
  fi
  
  echo "$SHA256  $OUTPUT_NAME.$ARCHIVE_EXT" >> "$OUTPUT_DIR/$CHECKSUM_FILE"
  
  rm -rf "$OUTPUT_DIR/$OUTPUT_NAME"
done

echo '```' >> "$OUTPUT_DIR/$CHECKSUM_FILE"

echo "✅ All builds completed and SHA256 checksums generated in $OUTPUT_DIR/$CHECKSUM_FILE"
