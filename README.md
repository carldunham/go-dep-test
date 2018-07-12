
Basically, this follows a pattern for a work project that has a few constraints, but I don’t think they are unusual or unexpected:
1. A few apps (device, cloud, etc) doing things independently, but as part of one overall project, and sharing some project-specific packages
2. Using dep for dependency management (that can flex)
3. Other languages besides Go (so I’m limited in what I can drop into the top-level `src/`, but can put what I want in my per-app dirs, and Go-specific `libs/`)
4. At some point soon, an overall build/CI process will run against the whole repo to build the apps (Go and non-Go)

The primary issue is that the shared dependencies aren’t recognized as being the same package (and exporting identical types) because they get interpreted as being from their respective `vendor` dirs.

Some options I’ve considered, that either have not worked, or I don’t know how to do:
- Don’t vendor the libs. This is what the general recommendation seems to be, but I can’t get them to build without their dependencies (and I want to run unit tests, etc)
- Build the libs as binaries that link into the apps. No idea how to do this
- Treat the libs as external dependencies, even tho they are in the same repo. I tried some naive approaches to including them in the apps’ `Gopkg.toml`, to no avail.

Some other points: the example has just one app, but the real project has three so far. It uses Ginkgo for unit testing, which sets flags in it’s `init()`, so yeah, that’s a problem. However, I assume that the solution to getting `github.com/google/uuid` to work will generalize. I chose that for the example, because Google.
