# Report Generator

![](./fig/architecture.drawio.png)

日報から簡単にレポートを出力できるインテグレーション。

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
- -r （教室名）
- -t （テンプレート）
- -om （markdown 出力名）
- -oe （excel 出力名）

## Installation

1. [Google Form(Template)](https://docs.google.com/forms/d/1Ee8WxMvLLZvv-1_nwWR74_5q8O2jwHLtWaV7zgBr7tg/edit?usp=sharing) を用意
2. Google Form の回答タブから Google Spreadsheet  を選択 スプレッドシートに移動。
3. csv 形式でダウンロード。
4. プログラムを実行。