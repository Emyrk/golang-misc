module github.com/Emyrk/golang-misc/pegnetcharity

go 1.13

replace github.com/Factom-Asset-Tokens/factom => github.com/Emyrk/factom v0.0.0-20191121205237-77d7cc617344

replace crawshaw.io/sqlite => github.com/AdamSLevy/sqlite v0.1.3-0.20191014215059-b98bb18889de

replace github.com/spf13/pflag v1.0.3 => github.com/AdamSLevy/pflag v1.0.4

require (
	github.com/Factom-Asset-Tokens/factom v0.0.0-20191120022136-7bf60a31a324
	github.com/FactomWyomingEntity/prosper-pool v0.4.0
	github.com/pegnet/pegnetd v0.2.1-0.20191206161753-99b44555aff8
)
