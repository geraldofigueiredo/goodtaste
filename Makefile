.PHONY: mock-gen
mock-gen:
	mockery --case underscore --all --dir services/web/front --output services/web/front/pkg/mocks