DAY=$(shell find . -type d -name "day*" | sed 's|./day||' | sort -n | tail -n 1)
NEXT_DAY=$(shell echo $$(($(DAY)+1)))

.PHONY: new

new:
	@mkdir -p day$(NEXT_DAY)/part1 day$(NEXT_DAY)/part2
	@touch day$(NEXT_DAY)/part1/main.go day$(NEXT_DAY)/part2/main.go day$(NEXT_DAY)/input.txt
	@echo "Created directories and files for day$(NEXT_DAY)"