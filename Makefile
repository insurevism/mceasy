.PHONY: clean schema mocks wire test build migration run

WIRE_DIR := internal/applications
OPENAPI_ENTRY_POINT := cmd/main.go
OPENAPI_OUTPUT_DIR := cmd/docs

ENT_PATH=./ent/schema
FEATURES=intercept,sql/modifier,sql/execquery

WIRE_ALL_DIR = /internal/applications

# Build the project
build:
	go build -o main cmd/main.go

# Run tests and generate coverage report
run:
	./main

# Run tests and generate coverage report
test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Generate ent models
schema:
	go generate ./ent

schema-advance:
	@echo "*****\n Advance mode will granted you a super power, use it wisely\n [Generate with entgo feature sql/modifier,sql/execquery]\n*****"
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature $(FEATURES) $(ENT_PATH)
	@echo "Running er_ps patch script - replace master_product_er_ps with master_product_erps..."
	@chmod +x ./er_ps_patch.sh
	@./er_ps_patch.sh
	@echo "Patch completed!"

# List all available schemas
schema-list:
	@echo "*****\n Available schemas:\n*****"
	@ls -1 $(ENT_PATH)/*.go | xargs -n 1 basename | sed 's/\.go//'

# Generate mockery mocks
mocks:
	mockery --all --dir internal --output mocks --packageprefix mock_ --keeptree

wire:
	@echo "Enter directory: "; \
	read dir; \
	echo "Accessing directory and wire all DI $(WIRE_DIR)/$$dir"; \
	cd $(WIRE_DIR)/$$dir && wire

wire-all:
	@success=0; \
	total=0; \
	skipped=0; \
	for dir in $$(ls internal/applications); do \
		if [ $$(find "internal/applications/$$dir" -name "*_injector.go" | wc -l) -gt 0 ]; then \
			total=$$((total + 1)); \
			echo "Wiring directory: $$dir"; \
			if (cd internal/applications/$$dir && wire && cd ../..); then \
				success=$$((success + 1)); \
				echo "✅ Successfully wired $$dir"; \
			else \
				echo "❌ Error while wiring $$dir. Stopping process."; \
				echo "Summary: $$success/$$total directories successfully wired ($$skipped skipped)"; \
				exit 1; \
			fi; \
		else \
			skipped=$$((skipped + 1)); \
			echo "!!! Skipping $$dir (no *_injector.go file found) !!!"; \
		fi; \
	done; \
	echo "Summary: $$success/$$total directories successfully wired ($$skipped skipped)"

# Generate OpenAPI Docs
swagger:
	swag fmt && swag init -g $(OPENAPI_ENTRY_POINT) -o $(OPENAPI_OUTPUT_DIR)

confirm:
	@read -p "$(shell echo -e '\033[0;31m') Warning: This action will clean up coverage reports, ent schema, and mockery generated codes. Do you want to continue? [Y/n]: $(shell tput sgr0)" choice; \
	if [ "$$choice" != "Y" ]; then \
		echo "$(shell echo -e '\033[0;31m') Terminating the clean-up process.$(shell output sgr0)"; \
    	exit 1; \
    fi

clean: confirm
	@echo "Warning this action will clean-up coverage report, ent schema and mockery generated codes "
	sleep 10

	@echo "Deleting coverage.out coverage.html on 5s"
	sleep 5
	rm -f coverage.out coverage.html

	@echo "Deleting all directories and files in ./ent except ./ent/schema and ./ent/generate.go on 5s"
	sleep 5
	@find ./ent/* ! -path "./ent/schema*" ! -path "./ent/generate.go" ! -path "./ent/hook*" -delete

	@echo "Deleting all directories and files in /.mocks on 5s"
	sleep 5
	rm -rf ./mocks/*

all: schema mocks swagger test build run

migration-build:
	@echo "Warning this action will build unix executable file "
	go build -v -o migration migrations/cmd/main.go

migration-create:
	@if [ ! -f "./migration" ]; then \
		$(MAKE) migration-build; \
    fi
	./migration mysql create $(name) $(type)

migration-up:
	@if [ ! -f "./migration" ]; then \
		$(MAKE) migration-build; \
    fi
	./migration mysql up

migration-down:
	@if [ ! -f "./migration" ]; then \
		$(MAKE) migration-build; \
    fi
	./migration mysql down

migration-down-to:
	@if [ ! -f "./migration" ]; then \
		$(MAKE) migration-build; \
    fi
	@if [ -z "$(version)" ]; then \
		echo "Error: Version cannot be empty."; \
		exit 1; \
	elif ! [[ "$(version)" =~ ^[0-9]+$$ ]]; then \
		echo "Error: Version must be a positive integer."; \
		exit 1; \
	elif [ "$(version)" = "0" ]; then \
		echo "Error: Version 0 is not allowed."; \
		exit 1; \
	elif ./migration mysql status | grep -q "Current: $(version)"; then \
		echo "Error: Version $(version) is the current version or higher."; \
		exit 1; \
	else \
		echo "Warning: This action will rollback to version $(version)"; \
		go build -v -o migration migrations/cmd/main.go; \
		./migration mysql down-to $(version); \
	fi

migration-status:
	@if [ ! -f "./migration" ]; then \
  		$(MAKE) migration-build; \
    fi
	./migration mysql status

patch-erps:
	@echo "Running er_ps patch script - replace master_product_er_ps with master_product_erps..."
	@chmod +x ./er_ps_patch.sh
	@./er_ps_patch.sh
	@echo "Patch completed!"