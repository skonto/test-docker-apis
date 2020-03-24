#!/bin/bash


# Used from https://ops.tips/blog/inspecting-docker-image-without-pull/

# Retrieve the digest, now specifying in the header
# the token we just received.
# note.:        $token corresponds to the token
#               that we received from the `get_token`
#               call.
# note.:        $image must be the full image name without
#               the registry part, e.g.: `nginx` should
#               be named `library/nginx`.
get_digest() {
  local image=$1
  local tag=$2
  local token=$3

  echo "Retrieving image digest.
    IMAGE:  $image
    TAG:    $tag
    TOKEN:  $token
  " >&2

  curl \
    --silent \
    --header "Accept: application/vnd.docker.distribution.manifest.v2+json" \
    --header "Authorization: Bearer $token" \
    "https://registry-1.docker.io/v2/$image/manifests/$tag" \
    | jq -r '.config.digest'
}

set -o errexit

main() {
  check_args "$@"

  local image=$1
  local tag=$2
  local token=$(get_token $image)
  local digest=$(get_digest $image $tag $token)

  start=`date +%s`
  get_image_configuration $image $token $digest
  end=`date +%s`

  echo "get image config:$((end-start))"

  search_image $image $token $digest
}

get_image_configuration() {
  local image=$1
  local token=$2
  local digest=$3

  echo "Retrieving Image Configuration.
    IMAGE:  $image
    TOKEN:  $token
    DIGEST: $digest
  " >&2

  curlS \
    --silent \
    --location \
    --header "Authorization: Bearer $token" \
    "https://registry-1.docker.io/v2/$image/blobs/$digest" \
    | jq -r '.'
}

get_token() {
  local image=$1

  echo "Retrieving Docker Hub token.
    IMAGE: $image
  " >&2

  curl \
    --silent \
    "https://auth.docker.io/token?scope=repository:$image:pull&service=registry.docker.io" \
    | jq -r '.token'
}

# Retrieve the digest, now specifying in the header
# that we have a token (so we can pe...
get_digest() {
  local image=$1
  local tag=$2
  local token=$3

  echo "Retrieving image digest.
    IMAGE:  $image
    TAG:    $tag
    TOKEN:  $token
  " >&2

  curl \
    --silent \
    --header "Accept: application/vnd.docker.distribution.manifest.v2+json" \
    --header "Authorization: Bearer $token" \
    "https://registry-1.docker.io/v2/$image/manifests/$tag" \
    | jq -r '.config.digest'
}

check_args() {
  if (($# != 2)); then
    echo "Error:
    Two arguments must be provided - $# provided.

    Usage:
      docker <image> <tag>

Aborting."
    exit 1
  fi
}

main "$@"
