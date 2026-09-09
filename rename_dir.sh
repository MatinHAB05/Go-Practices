#!/bin/sh
set -e
digits=3

for dir in [0-9]*; do
  if [ -d "$dir" ]; then
    num=$(echo "$dir" | grep -o '^[0-9]\+')
    rest=$(echo "$dir" | sed 's/^[0-9]\+//')
    newnum=$(printf "%0${digits}d" "$num")
    newname="${newnum}${rest}"

    if [ "$dir" != "$newname" ]; then
      if [ -e "$newname" ]; then
        echo "Skipped: '$newname' already exists → '$dir'"
      else
        echo "Renaming: '$dir' → '$newname'"
        mv -- "$dir" "$newname"
      fi
    fi
  fi
done