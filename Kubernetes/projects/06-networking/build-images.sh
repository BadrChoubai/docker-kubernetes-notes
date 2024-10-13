
(
  pushd users-api || exit
  docker build -t cc-users-api:1.0 .
  popd || exit
)
