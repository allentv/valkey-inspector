#!/bin/sh

echo "Initializing Valkey with sample data..."

# Strings
valkey-cli SET user:100 "Alice"
valkey-cli SET user:101 "Bob"

# Hashes
valkey-cli HSET config:app theme "dark" version "1.0.0" max_connections "50"

# Lists
valkey-cli LPUSH logs:errors "Error 1" "Error 2" "Critical Failure"

# Sets
valkey-cli SADD roles:admin "user:100" "user:200"

# Sorted Sets
valkey-cli ZADD leaderboard 100 "PlayerOne" 200 "PlayerTwo" 50 "PlayerThree"

echo "Initialization complete."