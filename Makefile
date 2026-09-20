DAYS := 01 02 03 04 05 06 07 08 09 10 11 12

.PHONY: run test

run:
	@for day in $(DAYS); do \
		echo "Dia $$day"; \
		go run ./day$$day; \
	done

test:
	go test ./...
