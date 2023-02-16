#!/bin/sh

if [ -z "$GITHUB_USERNAME" ]
then
  echo '$GITHUB_USERNAME is not set, trying to load .env file'
  . ./.env
fi


links=$(curl "https://github.com/$GITHUB_USERNAME/$GITHUB_USERNAME/blob/main/README.md" 2>/dev/null | grep -o https://camo.githubusercontent.com[0-9a-z\/]*)
for link in "${links[@]}"
do
   curl -X PURGE $link
done


