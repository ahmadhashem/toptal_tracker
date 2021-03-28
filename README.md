# This project is simple automation for generating report from toptal freelancer time tracking app https://tracker.toptal.com/

### I used simple bash and curl to retrieve the repot + using jq to parse json response from toptal
### Then go to send email using gopkg.in/mail.v2 to use smtp.

For simplicity I assumed you log time against one project always on toptal tracking app
Before you start, please replace your values in the file vars.env.
You need to replace values for:
- toptal username/password for toptal
- from/to mail users that would send/receive the email
- smtp username/password

Ideally this could be placed within a daily/monthly cron to send share the report


------------


Then to start, simply run:
``
docker-compose up -V --build
``
