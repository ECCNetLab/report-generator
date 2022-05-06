# Report Generator

## Syntax

```
go run ./main.go [-i <csv> -c <club name> -t <excel path>]

example:
go run ./main.go \
  -c ネットワーク研究会 \
  -r 2301 \
  -i file.csv \
  -t ./月間活動報告書.xlsx \
  -i ./ECCコンピュータ専門学校\ ネットワーク研究会\ .csv
```

## Parameters

- -i （入力する csv）
- -c （クラブ名）
- -t （テンプレート）
- -om （markdown 出力名）
- -oe （excel 出力名）

## Installation

1. Google Form の回答タブから `回答をダウンロード(.csv)` を選択
2. 実行