module github.com/krok-o/krokctl

go 1.16

require (
	github.com/krok-o/krok v0.0.0-20210610053306-6d5b590f1001
	github.com/olekukonko/tablewriter v0.0.5
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
)

replace (
	github.com/krok-o/krok => ../krok
)