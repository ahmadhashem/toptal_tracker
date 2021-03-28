#!/bin/bash

go get gopkg.in/mail.v2

apt update

apt install jq -y

START_DATE=$(date "+%F" -d "$(date +'%Y-%m-01')")
END_DATE=$(date +%F)

RESPONSE_JSON=$(curl -H "Content-Type: application/json" https://tracker-api.toptal.com/sessions --data "{\"email\": \"$TOPTAL_USERNAME\", \"password\": \"$TOPTAL_PASSWORD\", \"remember_me\": false}")

TOKEN=$(echo $RESPONSE_JSON | jq -r .access_token)
WORKER_ID=$(echo $RESPONSE_JSON | jq -r .profiles[].id)
PROJECT_ID=$(curl -G -H "content-type: application/json" "https://tracker-api.toptal.com/web/projects?archived=true&access_token=$TOKEN" | jq .projects[].id)
REPORT_NAME="$START_DATE-$END_DATE"

curl -G -o "$REPORT_NAME.csv" "https://tracker-api.toptal.com/reports/summary_csv?project_ids[]=$PROJECT_ID&worker_ids[]=$WORKER_ID&start_date=$START_DATE&end_date=$END_DATE&access_token=$TOKEN"

go run /sendmail.go "Monthly toptal timesheet report $REPORT_NAME" "PFA Report for current month" "$REPORT_NAME.csv"
