package preview

import (
"bytes"
"context"
"fmt"
"image"
"image/jpeg"
"os"

"github.com/gtsteffaniak/filebrowser/backend/pkg/indexing/iteminfo"
_ "github.com/oov/psd"
)

// generatePSDPreview decodes a PSD file's flattened/merged image and
// returns it encoded as full-size JPEG bytes. The standard resize
// pipeline scales it down to the requested preview size afterwards.
func (s *Service) generatePSDPreview(ctx context.Context, file iteminfo.ExtendedFileInfo, previewSize string) ([]byte, error) {
f, err := os.Open(file.RealPath)
if err != nil {
return nil, fmt.Errorf("failed to open PSD file: %w", err)
}
defer f.Close()

img, _, err := image.Decode(f)
if err != nil {
return nil, fmt.Errorf("failed to decode PSD image: %w", err)
}

var buf bytes.Buffer
if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
return nil, fmt.Errorf("failed to encode PSD preview: %w", err)
}
return buf.Bytes(), nil
}
