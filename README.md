# WebApp 2024

## プロジェクト概要
WebApp 2024 は、GoとGinフレームワークを使用して開発された軽量なWebアプリケーションです。MVCアーキテクチャを採用し、効率的かつモジュール化された開発をサポートします。

---
## 機能
- Ginフレームワークを使用した軽量なWebアプリケーション。
- 静的ファイルと動的HTMLテンプレートのレンダリングに対応。
- モジュール化されたMVCデザイン。
- 拡張可能な設定管理。

---

## 技術スタック
- **バックエンド**: Go (Ginフレームワーク)
- **フロントエンド**: HTML、CSS、JavaScript
- **テンプレートエンジン**: GinのHTMLテンプレートレンダリング
- **ビルドツール**: Goの標準ビルド機能

---

## 使用方法

### 1. コードをクローンする
```bash
git clone https://github.com/Michelinchef/webapp2024.git
cd webapp2024
```

### 2. データベースの設定
アプリケーションを実行する前に、データベースに接続する必要があります。

以下の手順に従って、デフォルトのデータベースを設定してください：

1. **MySQLをインストール**（未インストールの場合）。
2. **新しいデータベースを作成**：
   データベース名を `test` としてください。
   ```sql
   CREATE DATABASE test;

3. **MySQLをインストール**（dockerの場合、おすすめ）。
   ```
   docker pull mysql:latest
   docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=123456 -P 3306:3306 -d mysql: latest
   docker exec -it mysql-container bash
   mysql -u root -p
   CREATE DATABASE test;
   docker ps
   ```

### 3. 接続情報：

    ユーザー名：root
    パスワード：123456
    接続先：127.0.0.1
    ポート番号：3306

デフォルトの接続URL：
root:123456@tcp(127.0.0.1:3306)/test?


事前ビルド済みファイル

build フォルダには、以下のプラットフォーム向けにビルドされた実行ファイルが用意されています：

    blogapp-windows-amd64.exe: Windows用
    blogapp-linux-amd64: Linux用
    blogapp-macos-arm64: macOS用

使用方法: 各プラットフォームに対応するファイルをプロジェクトのルートディレクトリにコピーし、実行してください。 例:

./blogapp-linux-amd64
