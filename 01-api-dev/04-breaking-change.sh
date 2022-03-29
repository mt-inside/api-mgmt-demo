echo "No local changes?"
read
git checkout breaking-change
buf breaking --against ../../.git#branch=main,subdir=proto/example
git checkout main
