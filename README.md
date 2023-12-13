# Notify-Filterbox

![golang](https://img.shields.io/github/actions/workflow/status/starudream/notify-filterbox/golang.yml?style=for-the-badge&logo=github&label=golang)
![release](https://img.shields.io/github/v/release/starudream/notify-filterbox?style=for-the-badge)
![license](https://img.shields.io/github/license/starudream/notify-filterbox?style=for-the-badge)

## Precondition

- `Android`
    - [filterbox](https://filterbox.catchingnow.com)
    - [google play](https://play.google.com/store/apps/details?id=com.catchingnow.np)

```
{
  "android.title": "{android.title}",
  "android.text": "{android.text}",
  "android.bigText": "{android.bigText}",
  "filterbox.field.APP_NAME": "{filterbox.field.APP_NAME}",
  "filterbox.field.PACKAGE_NAME": "{filterbox.field.PACKAGE_NAME}",
  "filterbox.field.CHANNEL_ID": "{filterbox.field.CHANNEL_ID}",
  "filterbox.field.ONGOING": {filterbox.field.ONGOING}
  "filterbox.field.WHEN": {filterbox.field.WHEN}
}
```

- `MacOS`
  - [alerter](https://github.com/vjeantet/alerter)

```shell
brew install alerter
```

## Usage

```
Usage:
  notify-filterbox [flags]
  notify-filterbox [command]

Available Commands:
  service     Manage service

Flags:
      --addr string     server address (default "0.0.0.0:9781")
  -c, --config string   path to config file
  -h, --help            help for notify-filterbox
  -v, --version         version for notify-filterbox

Use "notify-filterbox [command] --help" for more information about a command.
```

## [License](./LICENSE)
