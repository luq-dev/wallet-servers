#! /bin/bash
clear

echo "Running Migrations..."
psql -U dl green_test -f migrations/*.sql 
echo ""

go run main.go