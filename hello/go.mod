module go-first/hello

go 1.22.1

replace go-first/func1 => ../func1
require go-first/func1 v0.0.0-00010101000000-000000000000

replace go-first => ../../go-first
require go-first v0.0.0