if ! which gifsicle ; then
  echo "gifsicle was not found, exiting"
  return 1
fi

if [ -z "$TARGET_LOCATION" ]
then
  echo '$TARGET_LOCATION is not set, trying to load .env file'
  . ./.env
fi

gifsicle -i $TARGET_LOCATION -O3 -o $TARGET_LOCATION
