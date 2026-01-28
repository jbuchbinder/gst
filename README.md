### Go bindings for GStreamer at a very early stage of maturity.

This package is based on [GLib bindings](https://github.com/falinux/glib). It
should be goinstalable. Try

```
$ go get github.com/jbuchbinder/gst
```

If it does not compile on an x86_64 machine, try this first:

```
export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig
```

For Ubuntu 24.04+ you'll need to make sure you install:

 * libgtk2.0-dev
 * libgio-2.0-dev
 * libgstreamer1.0-dev
 * libgstreamer-plugins-base1.0-dev

#### Documentation

See *examples* directory and http://gopkgdoc.appspot.com/pkg/github.com/falinux/gst

To run examples use `go run` command:

```
$ cd examples
$ go run simple.go
```

To run live WebM example use `go run live_webm.go` and open
http://127.0.0.1:8080 with your browser. You probably need to wait a long time
for video because of small bitrate of this stream and big buffer in you browser.
