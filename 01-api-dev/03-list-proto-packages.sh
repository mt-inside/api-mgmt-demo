buf build --exclude-source-info -o -#format=json | jq '.file[] | .package' | sort | uniq
