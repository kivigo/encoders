module github.com/kivigo/encoders/testdata

go 1.25.1

replace github.com/kivigo/encoders/encrypt => ../encrypt

require github.com/kivigo/encoders/encrypt v0.0.0-00010101000000-000000000000

require (
	github.com/azrod/cryptio v1.0.0 // indirect
	github.com/kivigo/encoders v0.0.0-20250914204157-1fac3f1828ac // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
)
