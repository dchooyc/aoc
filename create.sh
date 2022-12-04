#!/bin/bash
# This creates a new puzzle folder with the template

# Name folder
read -p "Please input puzzle folder name: " puzzle

# Make folder
mkdir $puzzle

# Copy template files
cp "template/"* $puzzle

echo "The new puzzle folder $puzzle has been created!"