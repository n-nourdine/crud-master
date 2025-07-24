#!/bin/bash
sudo apt-get install -y postgresql postgresql-contrib
sudo -u postgres psql -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"
sudo -u postgres createdb $MOVIES_DB
sudo -u postgres createdb $BILLING_DB
sudo -u postgres psql $MOVIES_DB -c "CREATE TABLE movies (id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, description TEXT);"
sudo -u postgres psql $BILLING_DB -c "CREATE TABLE orders (id SERIAL PRIMARY KEY, user_id VARCHAR(50), number_of_items INT, total_amount DECIMAL);"