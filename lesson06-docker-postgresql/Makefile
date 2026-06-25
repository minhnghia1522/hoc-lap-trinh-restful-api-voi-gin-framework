importdb:
	docker exec -i postgres-db-go psql -U root -d master-golang < ./backupdb-master-golang.sql
exportdb:
	docker exec -i postgres-db-go pg_dump -U root -d master-golang > ./backupdb-master-golang.sql