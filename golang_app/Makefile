# Start containers
docker-up: 
	docker compose up -d

# New ent schema
new: 
	go run entgo.io/ent/cmd/ent new 

# Generate ent code from schemas
generate:
	go generate ./ent

# Migrate Ent scheme to database
migrate migration:
	go run database/migrate/main.go

# Seed database with data
seed:
	go run database/seed/main.go

# Server start
start:
	air

# Avoid file name conflicts
.PHONY:

