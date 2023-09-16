PWD = $(CURDIR)
# Test timeout
TEST_TIMEOUT?=20s

# creates coverage report
.PHONY: cover
cover:
	go test -timeout=$(TEST_TIMEOUT) -v -coverprofile=coverage.out ./...  && go tool cover -html=coverage.out

# creates coverage report
.PHONY: test
test:
	go test -timeout=$(TEST_TIMEOUT) -v ./...
