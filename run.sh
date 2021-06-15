#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings -dbname=bookings -dbuser=avizin -production=false -cache=false