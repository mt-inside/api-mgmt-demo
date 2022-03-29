echo "No local changes?"
read
git checkout breaking-change
buf breaking --against .git#branch=main
git checkout main
