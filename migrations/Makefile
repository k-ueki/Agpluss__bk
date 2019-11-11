DBNAME:=agpluss
DATASOURCE:=root:@tcp(127.0.0.1:3306)/$(DBNAME)

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DBNAME)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" 

migrate/status:
	@goose -dir migrations mysql "$(DATASOURCE)" status

migrate/up:
	@goose -dir migrations mysql "$(DATASOURCE)" up

__migrate/down:
	@goose -dir migrations mysql "$(DATASOURCE)" down
