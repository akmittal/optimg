# Optimg
Optimizes all the images at path to modern formats like AVIF, WEBP and to optimized jpeg. Server most optimized format supported by client browser to reduce data transfer and faster client page loads.


## Prerequisites

- [libvips](https://github.com/libvips/libvips) 8.3+ (8.8+ recommended)
- C compatible compiler such as gcc 4.6+ or clang 3.0+
- Go 1.3+

**Note**: 
 * `libvips` v8.3+ is required for GIF, PDF and SVG support.
 * `libvips` v8.9+ is required for AVIF support. `libheif` compiled with a AVIF en-/decoder also needs to be present.

## Installation

```bash
go install github.com/akmittal/optimg/server/cmd/optimg
```

### libvips

Follow `libvips` installation instructions:

[https://libvips.github.io/libvips/install.html](https://libvips.github.io/libvips/install.html)


## Command line usage
### 1. Convert images 
``` text
Usage:
  optimg optimize [flags]

Flags:
  -f, --format strings      Image format
  -h, --help                help for optimize
  -q, --quality ints        Image quality
  -s, --sourcePath string   Source Image Path (default "/tmp")
  -t, --targetPath string   Image Path where to store images (default "/tmp")

required flag(s) "format", "quality", "sourcePath", "targetPath" 
```
Example: 
``` bash
optimg optimize -s /tmp/source -t /tmp/dest -f webp -q 50 -f jpg -q 80
```

### Server images
Image server will auto serve best available format for image. If request is for `/earth.jpeg` and `/earth.avif` is available then `/earth.avif` will be served if supported by browser
``` text
Usage:
  optimg serve [flags]

Flags:
  -d, --directory string   Directory which is to be served (default "/tmp")
  -h, --help               help for serve
  -a, --host string        Application host where to run (default "localhost")
  -p, --port string        Optimg dashboard port (default "8000")

required flag(s) "directory"
```

Images will be served on `http://host:port/public/{imagepath}`