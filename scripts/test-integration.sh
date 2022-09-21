#!/bin/bash


BASE_URL="http://127.0.0.1:8080"

# Function that call using either POST or PUT and report errors
function call_put_post {
  local url=$1
  local method=$2
  local file_body=$3
  local expected=$4
  local response=$(curl --request $method -s $url -w "%{http_code}" --data-binary @$file_body)
  local body=${response::-3}
  local status=$(printf "%s" "$response" | tail -c 3)
  
  if [ "$status" -ne $expected ]; then
        echo "ERROR: HTTP $method repsonse is $status, error $body"
        exit
  fi
  echo "Call $method $url Response: $body"

}

# Function that call using GET and report errors
function call_get {
  local url=$1
  local response=$(curl -s $url -w "%{http_code}")
  local body=${response::-3}
  local status=$(printf "%s" "$response" | tail -c 3)
  
  if [ "$status" -ne $2 ]; then
        echo "ERROR: HTTP repsonse is $status, error $body"
        exit
  fi
  echo "Call GET $url Response: $body"

}

# Test cases
call_get "$BASE_URL/20200201/" "200"
# Get , Create , Update and delete 
call_get "$BASE_URL/20200201/cars" "200"
call_get "$BASE_URL/20200201/cars/GLD-134" "404"
call_put_post "$BASE_URL/20200201/cars" "POST" "./body.json" "200"
call_get "$BASE_URL/20200201/cars/GLD-CA01" "200"
call_put_post "$BASE_URL/20200201/cars" "PUT" "./body-update.json" "200"
call_get "$BASE_URL/20200201/cars/GLD-CA01" "200"
# List
call_get "$BASE_URL/20200201/cars" "200"