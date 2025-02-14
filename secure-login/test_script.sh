#!/bin/bash

BASE_URL="http://localhost:8080"

# Test Register
echo "Testing Register Endpoint"
curl -X POST -d "username=zakaria7&password=Password7" "$BASE_URL/register"
echo -e "\n"

# Test Login
echo "Testing Login Endpoint"
curl -X POST -d "username=zakaria7&password=Password7" -c cookies.txt "$BASE_URL/login"
echo -e "\n"

# Test Private Access with CSRF token
echo "Testing Private Endpoint"
CSRF_TOKEN=$(awk '/csrf_token/ {print $7}' cookies.txt)
curl -X POST -d "username=zakaria7" -b cookies.txt -H "X-CSRF-Token: $CSRF_TOKEN" "$BASE_URL/private"
echo -e "\n"

# Test Logout
echo "Testing Logout Endpoint"
curl -X POST -d "username=zakaria7" -b cookies.txt "$BASE_URL/logout"
echo -e "\n"
