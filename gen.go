package main

//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-0-4.json --output-file src/types/oscal-1-0-4/types.go --tags json,yaml --pkg oscalTypes_1_0_4
//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-0-5.json --output-file src/types/oscal-1-0-5/types.go --tags json,yaml --pkg oscalTypes_1_0_5
//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-0-6.json --output-file src/types/oscal-1-0-6/types.go --tags json,yaml --pkg oscalTypes_1_0_6
//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-1-0.json --output-file src/types/oscal-1-1-0/types.go --tags json,yaml --pkg oscalTypes_1_1_0
//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-1-1.json --output-file src/types/oscal-1-1-1/types.go --tags json,yaml --pkg oscalTypes_1_1_1
//go:generate go run main.go generate generate --input-file=src/internal/schemas/oscal_complete_schema-1-1-2.json --output-file src/types/oscal-1-1-2/types.go --tags json,yaml --pkg oscalTypes_1_1_2
