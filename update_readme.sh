#!/bin/bash

# Extract the day numbers from the directory names
progress=$(find . -name "day_*" | cut -d _ -f 2 | sort -n)

# Read the current README content
readme_content=$(cat README.md)

# Find the start and end of the Progress section
progress_start=$(echo "$readme_content" | grep -n "## Progress" | cut -d: -f1)

# Extract the content before and after the Progress section
before_progress=$(echo "$readme_content" | head -n $((progress_start)))

# Generate the new Progress section
new_progress=""
for day in $progress; do
    new_progress="$new_progress- Day $day\n"
done

# Combine the content before, the new Progress section, and the content after
new_readme_content="$before_progress\n$new_progress"

# Write the updated content back to the README file
echo -e "$new_readme_content" > readme.md

echo "README updated with the latest progress."
