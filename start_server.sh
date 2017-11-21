APP_ENV="${REVEL_APP_ENV:-"PROD"}"
if [ ${APP_ENV} = "PROD" ]; then
  echo "hello, PROD"
  revel run near_me_server prod ${PORT}
else
  echo "hello, DEV"
  revel run near_me_server
fi
