## If project doesn't update when running:
go clean -cache -modcache -i -r
go install .
gopherflow generate

## TODO:
restrict user input while generating
add clipboard copying
(add Git connection for intelligent commit messaging)