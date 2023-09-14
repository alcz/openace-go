module testdb

go 1.15

require (
	github.com/alcz/openace-go v0.0.0
	github.com/alcz/openace-go/oauto v0.0.0
	github.com/alcz/openace-go/onoerr v0.0.0
	github.com/alcz/openace-go/oerr v0.0.0
)

replace github.com/alcz/openace-go => ../
replace github.com/alcz/openace-go/oauto => ../oauto
replace github.com/alcz/openace-go/onoerr => ../onoerr
replace github.com/alcz/openace-go/oerr => ../oerr
