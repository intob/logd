# Generate a Go file with git blame context
# go_blame.sh

#!/bin/bash

OUTPUT_FILE="git_blame_context.go"
echo "package main" > $OUTPUT_FILE
echo "" >> $OUTPUT_FILE
echo "var GitBlameContext = map[string]map[int]string{" >> $OUTPUT_FILE

for file in $(find . -name '*.go'); do
    echo "  \"$(basename $file)\": {" >> $OUTPUT_FILE
    git blame --line-porcelain $file | grep -e "author " -e "author-mail " | awk '{
        if ($1 == "author") author=$2; if ($1 == "author-mail") print "    " NR/2 ": \"" author " <" $2 ">\""
    }' >> $OUTPUT_FILE
    echo "  }," >> $OUTPUT_FILE
done

echo "}" >> $OUTPUT_FILE
