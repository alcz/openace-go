# openace-go
## xBase database access library for Go lang
### these are wrappers for C library openace32.dll/.lib/.so based on Harbour

Go app can access/create DBFCDX/DBFNTX .dbf files on-disk, in-memory
or excersise advanced setups with NETIO file access on remote server,
Custom RDD LetoDBf can be available from openace too.

State of this library reflects my needs or that of my customers
and is provided here as-is without any warranty. Code in this
repository is available under MIT license.

### to use:
you have to get a binary build of openace32.dll - a rare beast ;)

and a compatible Harbour source tree (.h includes) on your drive
of course a Go (1.15+ tested) compiler is needed too

platforms other than 32-bit Windows are not tested currently,
but if you work out building/linking yourself it should be even easier
as no C 32-bit stdcall convention / binary compatibility with ace32.dll is needed

see build32.cmd preferably in examples/ 

### there are four different packages of functions in this lib:
- [openace.*](https://github.com/alcz/openace-go/goaddgen.txt) - nearly 1:1 mappings to C API, not convinient, reading a string needs manual buffer allocation
- [oauto.*](https://github.com/alcz/openace-go/oauto/aceauto.go)   - automatic string mappings, errors are redirected to Go panic()
- [oerr.*](https://github.com/alcz/openace-go/oerr/acewerr.go)    - automatic string mappings, errors returned as second value
- [onoerr.*](https://github.com/alcz/openace-go/onoerr/acenoerr.go)  - automatic string mappings, no errors are reported, only false
            or empty results are returned in such case. do not use for
            serious tasks, only for tests, quick-tools or with extreme care

### to regenerate wrappers:
also a copy of installed SWIG 4 is needed
plus a runnable Harbour compiler to run goaddgen.prg
see gen.cmd

### ideas and TODOs:
- move oauto.GetStringErr() and similar wrappers to oerr.GetString() naturally...
- allow access to fields by pos, ace32.dll achieves this by specifying low
  value in a pointer. Go does not have multiple overloads for functions,
  so we will eventually need separate set of GetStringByPos( wa, fieldpos )
- per workarea Go struct
- strongly typed workarea struct generator, with field checking on 
  database open
- generate C glue code without SWIG
- test other platforms
