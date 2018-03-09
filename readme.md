# Simple replace text tool

## Download

[windows-64](release/windows-64/replace.exe)

[linux-64](release/linux-64/replace)

[darwin-64](release/darwin-64/replace)

# Usage

```cmd

replace -c config.json

```

config.json

```json

{
  "replace":[{
    "paths":["test"],
    "regReplace":[
      {
        "regex":"yale8848",
        "replacement":"[yale8848]"
      }
    ],
    "recursive":true,
    "silent":false
  }]
}

```

