#!/bin/sh
set -e  # Exit immediately if a command exits with a non-zero status

# Run the API
./didlydoodash-migrate
./didlydoodash-api