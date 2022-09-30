#!/bin/bash 

if urlcheck -s http://go.dev -t 20 1>/dev/null 2>&1; then
  echo go.dev is UP
else
  echo go.dev is DOWN
fi