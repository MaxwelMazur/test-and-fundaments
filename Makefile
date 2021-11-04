Test: # exemplos 
	go test -v
	go test -cover
	go test -timeout 30s -run ^TestSumAll$ github.com/MaxwelMazur/test-and-fundaments/list
	go test -v -cover -timeout 30s  -run ^TestSumAll$