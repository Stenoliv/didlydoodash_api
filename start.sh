#!/bin/sh
set -e  # Exit immediately if a command exits with a non-zero status

# Flag file to indicate if migration has been run
FLAG_FILE="/app/migration_done"

# Check if the migration has already been run
if [ ! -f "$FLAG_FILE" ]; then
    echo "Running drop and migrate..."
    ./didlydoodash-drop
    ./didlydoodash-migrate

    # Create the flag file to indicate that migration has been run
    touch "$FLAG_FILE"
else
    echo "Migration has already been run. Skipping..."
fi

# Run the API
./didlydoodash-api