v0.3.0
======
Apr 24 2025

- 存在しないファイルを指定した時、エラーにならないという CLI の不具合を修正
- 指定したファイルから削除すべきファイル名を読むオプション `-from-file FILENAME` を追加 (標準入力の場合は `-` を使う)
- Windows API:`shFileOperationW` の結果もチェックし、`windows.ERROR_SUCCESS` を非エラーと扱うようにした。

v0.2.0
======
Nov 07 2023

- 非Windows で[the FreeDesktop.org Trash 1.0][freedesktop] の "the home trash" をサポート

[freedesktop]: https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html

v0.1.0
======
Nov 06 2023

- 初版: 非Windows では Trash は削除と同様に機能
