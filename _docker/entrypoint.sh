#!/bin/sh
set -e
if [ -d /srv ]; then
  chown filebrowser:filebrowser /srv 2>/dev/null || true
  mkdir -p /srv/.artbox 2>/dev/null || true
  chown filebrowser:filebrowser /srv/.artbox 2>/dev/null || true
fi
exec su-exec filebrowser ./filebrowser
