test:
	ab -p postdata.json -T application/json  -c 1 -n 10000 http://localhost:8080/gen/bl1
	ab -p postdata.json -T application/json  -c 1 -n 10000 http://localhost:8080/int/bl2

trigbl2:
	ab -p postdata.json -T application/json  -c 1 -n 1 http://localhost:8080/int/bl2

trigbl1:
	ab -p postdata.json -T application/json  -c 1 -n 1 http://localhost:8080/gen/bl1

run:
	go run .

.PHONY: pprof
pprof:
	go tool pprof -http=:8060  pprof

asm_generics:
	GOOS=linux GOARCH=amd64 go tool compile -S asm/generics/main.go

asm_ifcae:
	GOOS=linux GOARCH=amd64 go tool compile -S asm/iface/main.go