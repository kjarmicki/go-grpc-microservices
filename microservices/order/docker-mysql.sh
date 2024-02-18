#!/bin/bash

docker run -p 3300:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=order mysql
