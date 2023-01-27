# ting-task-nhk-easy [![Go](https://github.com/ting-app/ting-task-nhk-easy/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ting-app/ting-task-nhk-easy/actions/workflows/build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/ting-app/ting-task-nhk-easy)](https://goreportcard.com/report/github.com/ting-app/ting-task-nhk-easy)
A scheduled job that saves [NHK NEWS WEB EASY](https://www3.nhk.or.jp/news/easy/) as ting, the task runs at 10:30 AM (UTC) every day.

## Getting started
Run with docker:

```sh
docker run -e DB_USER_NAME=user name of MySQL database \
  -e DB_PASSWORD=password of MySQL user \
  -e DB_HOST=host of MySQL database \
  -e DB_PORT=port of MySQL database \
  -e ENABLE_SENTRY=true \
  -e SENTRY_DSN=your sentry dsn \
  -d xiaodanmao/ting-task-nhk-easy:latest
```

## License
[MIT](LICENSE)
