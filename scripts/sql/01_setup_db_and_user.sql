-- Filename: 01_setup_db_and_user.sql
-- Description: Creates the database and application user.
-- This script should be run by a PostgreSQL superuser.

-- Use \set to assign variables passed in from the psql command line
\set db_name `echo :DB_NAME`
\set db_user `echo :DB_USER`
\set db_pass `echo :DB_PASS`

-- Drop existing user and database to make the script rerunnable
DROP DATABASE IF EXISTS :db_name;
DROP ROLE IF EXISTS :db_user;

-- Create the user with a password
CREATE ROLE :db_user WITH LOGIN PASSWORD :'db_pass';

-- Create the database and set the new user as the owner
CREATE DATABASE :db_name WITH OWNER = :db_user;