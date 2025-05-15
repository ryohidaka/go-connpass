# Changelog

## [0.10.0](https://github.com/ryohidaka/go-connpass/compare/v0.9.0...v0.10.0) (2025-05-15)


### Features

* **api:** クエリパラメータを可変長引数に変更 ([9145be0](https://github.com/ryohidaka/go-connpass/commit/9145be07c79ff116e9469b89ec15a1d38e2ab06d))

## [0.9.0](https://github.com/ryohidaka/go-connpass/compare/v0.8.3...v0.9.0) (2025-05-11)


### Features

* **api:** APIキーの取得方法を変更 ([06d614f](https://github.com/ryohidaka/go-connpass/commit/06d614fec64ba1ab668bfb39b3925e6f0d3429ec))

## [0.8.3](https://github.com/ryohidaka/go-connpass/compare/v0.8.2...v0.8.3) (2025-05-06)


### Bug Fixes

* 非推奨のio/ioutilをioに変更 ([8d94b44](https://github.com/ryohidaka/go-connpass/commit/8d94b44d26e94f84a3ce7d36ff3c78b986d229b2))

## [0.8.2](https://github.com/ryohidaka/go-connpass/compare/v0.8.1...v0.8.2) (2025-05-06)


### Bug Fixes

* リンクの記載方法を修正 ([e96cd65](https://github.com/ryohidaka/go-connpass/commit/e96cd6525b5f017f4b42b218f0900a3720d26772))

## [0.8.1](https://github.com/ryohidaka/go-connpass/compare/v0.8.0...v0.8.1) (2025-05-05)


### Bug Fixes

* ファイル名をスネークケースに統一 ([f1c9222](https://github.com/ryohidaka/go-connpass/commit/f1c922244c8af88ddca323dd91de4facdef75296))

## [0.8.0](https://github.com/ryohidaka/go-connpass/compare/v0.7.0...v0.8.0) (2025-04-25)


### Features

* **api:** ユーザー発表イベント一覧取得を追加 ([cc91b83](https://github.com/ryohidaka/go-connpass/commit/cc91b83847c87800ab8e5da3efd6b046ab7a3b54))
* **types:** ユーザー発表イベント一覧の型定義を追加 ([34e604c](https://github.com/ryohidaka/go-connpass/commit/34e604cbf78b14a4882a09ad226915082f4485e2))

## [0.7.0](https://github.com/ryohidaka/go-connpass/compare/v0.6.0...v0.7.0) (2025-04-25)


### Features

* **api:** ユーザー参加イベント一覧取得を追加 ([7b27e7b](https://github.com/ryohidaka/go-connpass/commit/7b27e7b8cffc7268886c532e8aa46615e70854f2))
* **types:** ユーザー参加イベント一覧の型定義を追加 ([0143816](https://github.com/ryohidaka/go-connpass/commit/0143816463ce21fecfef1223dbbcb0c9ff1eb0b8))

## [0.6.0](https://github.com/ryohidaka/go-connpass/compare/v0.5.0...v0.6.0) (2025-04-25)


### Features

* **api:** ユーザー所属グループ一覧取得を追加 ([585ee90](https://github.com/ryohidaka/go-connpass/commit/585ee90acf1f11fe014a0a7223e046a4949dc6a9))
* **types:** ユーザー所属グループ一覧の型定義を追加 ([7544fad](https://github.com/ryohidaka/go-connpass/commit/7544fad4e1a3c49683f77241bfdb51ffde227893))

## [0.5.0](https://github.com/ryohidaka/go-connpass/compare/v0.4.2...v0.5.0) (2025-04-25)


### Features

* **api:** ユーザー一覧取得を追加 ([78d5dee](https://github.com/ryohidaka/go-connpass/commit/78d5deeac9fb8fb1e791b5d7dcc6ffcb627d29ba))
* **types:** ユーザー一覧の型定義を追加 ([90a4010](https://github.com/ryohidaka/go-connpass/commit/90a401054f8ecb94e339600f85f130e1692b9255))
* **types:** 型の名前を変更 ([543440f](https://github.com/ryohidaka/go-connpass/commit/543440f9c92229dcfa94ea01a945131fff234369))

## [0.4.2](https://github.com/ryohidaka/go-connpass/compare/v0.4.1...v0.4.2) (2025-04-25)


### Bug Fixes

* **types:** イベント一覧のパラメータを修正 ([d8ad6b6](https://github.com/ryohidaka/go-connpass/commit/d8ad6b66ae941c35938c6f0c1c3f953a254adbf6))

## [0.4.1](https://github.com/ryohidaka/go-connpass/compare/v0.4.0...v0.4.1) (2025-04-25)


### Bug Fixes

* **api:** クエリパラメータがポインタの場合に適用されない不具合を修正 ([2bd24fb](https://github.com/ryohidaka/go-connpass/commit/2bd24fb17f679231d3dad5fda1ce50a9f6836997))
* **types:** イベントIDのJSONのキー名を修正 ([e9dd109](https://github.com/ryohidaka/go-connpass/commit/e9dd109be5d2bc326d986e79d56ba73a4f3f9b67))
* **types:** イベント一覧取得時に指定するイベントIDが単一指定になっている不具合を修正 ([88aa31a](https://github.com/ryohidaka/go-connpass/commit/88aa31a5717b5f1f7c285bb5458fa469c1a6a4d6))
* **types:** イベント情報の構造体にイベント画像URLが含まれていない不具合を修正 ([39ae285](https://github.com/ryohidaka/go-connpass/commit/39ae2859b78e912b95d9820690b2ab3df21a8bed))

## [0.4.0](https://github.com/ryohidaka/go-connpass/compare/v0.3.0...v0.4.0) (2025-04-25)


### Features

* **api:** グループ一覧取得を追加 ([13bce86](https://github.com/ryohidaka/go-connpass/commit/13bce86cb5dad54d2c85a1e69fbab6da698b6718))
* **types:** グループ一覧の型定義を追加 ([1e312f3](https://github.com/ryohidaka/go-connpass/commit/1e312f33204d331441103ca63bf969be669390b9))

## [0.3.0](https://github.com/ryohidaka/go-connpass/compare/v0.2.0...v0.3.0) (2025-04-25)


### Features

* **api:** イベント資料一覧取得を追加 ([e133e3e](https://github.com/ryohidaka/go-connpass/commit/e133e3e95dcf0389bcb96d69aa189b4e7104ffe7))
* **types:** イベント資料一覧の型定義を追加 ([6d506d7](https://github.com/ryohidaka/go-connpass/commit/6d506d7e3e798f46d98173a831bd81730d3a61b6))

## [0.2.0](https://github.com/ryohidaka/go-connpass/compare/v0.1.1...v0.2.0) (2025-04-25)


### Features

* **api:** ベースURLとAPIバージョンを統合 ([90f8135](https://github.com/ryohidaka/go-connpass/commit/90f813564ba3bca672cc7d679ce4b52cbdac21f4))

## [0.1.1](https://github.com/ryohidaka/go-connpass/compare/v0.1.0...v0.1.1) (2025-04-24)


### Bug Fixes

* **api:** イベント検索用のクエリパラメータをオプショナルに修正 ([1a7821c](https://github.com/ryohidaka/go-connpass/commit/1a7821c3e17ddd53520d61055eab02fcd80b917d))

## 0.1.0 (2025-04-24)


### Features

* **api:** Connpass API にアクセスするためのクライアントを作成 ([e8750d1](https://github.com/ryohidaka/go-connpass/commit/e8750d18e8b66a3e15a4756391c0bf05c1a6572c))
* **api:** Goモジュールを初期化 ([f030fb0](https://github.com/ryohidaka/go-connpass/commit/f030fb09064097806705ecbaf18167d6525011a6))
* **api:** イベント一覧取得を追加 ([6eff2ee](https://github.com/ryohidaka/go-connpass/commit/6eff2ee121bcb37c57209b439a5bb5c9bd4c0c48))
* **api:** 共通のリクエスト処理を追加 ([2acab56](https://github.com/ryohidaka/go-connpass/commit/2acab563b2a9257ce7e7856401f2e2a9648ab54a))
* **types:** イベント一覧の型定義を追加 ([86f20bc](https://github.com/ryohidaka/go-connpass/commit/86f20bc8c35a535409f258daed58375143d58535))


### Miscellaneous Chores

* release 0.1.0 ([d7a7040](https://github.com/ryohidaka/go-connpass/commit/d7a7040a0f258cbb2f2fc05c2a3eef74eaf2a689))
