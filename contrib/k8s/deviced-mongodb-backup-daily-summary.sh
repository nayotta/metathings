#! /bin/bash

pip install -i pandas requests tabulate -i https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple

python3 /scripts/deviced-mongodb-backup-daily-summary.py
