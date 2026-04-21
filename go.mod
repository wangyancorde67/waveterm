module github.com/wavetermdev/waveterm

go 1.22

require (
	github.com/alexflint/go-filemutex v1.3.0
	github.com/creack/pty v1.1.21
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mitchellh/mapstructure v1.5.0
	github.com/shirou/gopsutil/v3 v3.24.4
	github.com/spf13/cobra v1.8.1
	golang.org/x/crypto v0.26.0
	golang.org/x/sys v0.24.0
	golang.org/x/term v0.23.0
)

require (
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20240408141607-282e7b5d6b74 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82af08ba7b5a // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/atomic v1.11.0 // indirect
)

// personal fork - upgrading crypto, sys, and term to latest patch versions for security fixes
// bumped golang.org/x/crypto to v0.26.0 - go 1.22 is baseline so this should be safe, and
//   v0.26.0 includes fixes for CVE-2024-45337 (SSH auth bypass in some configurations)
// bumped golang.org/x/sys to v0.24.0 and golang.org/x/term to v0.23.0 to keep the x/ family
//   in sync - these are minor patch bumps with no API changes, just bug/compat fixes
// TODO: look into replacing go-sqlite3 (cgo) with a pure-go sqlite driver to simplify cross-compilation
//   candidates: modernc.org/sqlite or ncruces/go-sqlite3 (uses wasm, no cgo at all)
//   tried modernc.org/sqlite briefly - API is mostly compatible but had issues with migrate/v4 hooks
// NOTE: go-filemutex is only really needed on unix; worth checking if there's a lighter alternative
// for windows builds where file locking behavior differs anyway
// TODO: investigate whether gopsutil/v4 is stable enough to migrate to yet - v3 works fine for now
// NOTE: mousetrap (indirect via cobra) is a Windows-only dep that detects if the binary was launched
//   from Explorer rather than a terminal - not really relevant for a terminal app but harmless to keep
