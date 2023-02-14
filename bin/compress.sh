if ! which gifsicle ; then
  echo "gifsicle was not found, exiting"
  return 1
fi

. ./.env
gifsicle -i $TARGET_LOCATION -O3 -o $TARGET_LOCATION
