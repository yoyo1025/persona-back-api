FROM postgres:15.1

# 必要なパッケージのインストール
RUN apt-get update && apt-get install -y locales

# ロケール設定
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8
