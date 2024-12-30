#! /usr/bin/env python3

import os
import time

import requests
import pandas as pd
from tabulate import tabulate

envBackupInprogLck = os.environ["BACKUP_INPROG_LOCK"]
envNewestMetadata = os.environ["NEWEST_METADATA"]
envSlackNotificationUrl = os.environ["SLACK_NOTIFICATION_URL"]

time.sleep(15)
while True:
    try:
        os.stat(envBackupInprogLck)
        print("backup is still progressing @" + pd.to_datetime(time.time()*1000000000+28800000000000).tz_localize('Asia/Shanghai').isoformat(timespec='milliseconds'), flush=True)
        time.sleep(3)
    except FileNotFoundError:
        break
time.sleep(3)

df = pd.read_csv(open(envNewestMetadata).read())
df['beginAt'] = pd.to_datetime(df.beginAt.apply(lambda x: x+'T00:00:00+08:00'))
df['endAt'] = pd.to_datetime(df.endAt.apply(lambda x: x+'T00:00:00+08:00'))
df['backupAt'] = pd.to_datetime(df.backupAt).apply(lambda x: x.tz_convert('Asia/Shanghai'))
df['size(megabytes)'] = df['size(bytes)']/1024/1024
df['compressedSize(megabytes)'] = df['compressedSize(bytes)']/1024/1024

beginAt=df.beginAt.iloc[0].strftime('%Y-%m-%d')
endAt=df.endAt.iloc[0].strftime('%Y-%m-%d')
jobStartAt=df.backupAt.min().isoformat(timespec='milliseconds')
jobEndAt=pd.to_datetime(df.backupAt.max().value+(df['elapsed(s)'].iloc[-1]*1000000000)).tz_localize('Asia/Shanghai').isoformat(timespec='milliseconds')
elapsedMinute=df['elapsed(s)'].sum()/60
documents=df['total'].sum()
documentsPerSecond=df['total'].sum()/86400
sizeMegabytes=df['size(bytes)'].sum()/1024/1024
avgDocumentSizeKilobytes=df['size(bytes)'].sum()/1024/df['total'].sum()
compressedSizeMegabytes=df['compressedSize(bytes)'].sum()/1024/1024

top10BiggestSizeDocumentsBlock = '''
Top 10 Biggest Size Doucments
'''

_top10BiggestSizeFlows = df.sort_values(by=['size(megabytes)'], ascending=False)
top10BiggestSizeDocumentsBlock += tabulate([(lambda data: [
        data['flowId'][:12], # Flow
        str(data['total']), # Documents
        data['size(megabytes)'], # Size(MB)
        data['compressedSize(megabytes)'], # CompressedSize(MB)
        data['backupAt'].isoformat(timespec='milliseconds'), # BackupAt
        data['elapsed(s)'], # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'], # BackupSpeed(MB/s)
        data['total']/1440, # DocumentSpeed(doc/min)
])(_top10BiggestSizeFlows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

top10BiggestTotalDocumentsBlock = '''
Top 10 Biggest Total Doucments
'''

_top10BiggestTotalFlows = df.sort_values(by=['total'], ascending=False)
top10BiggestTotalDocumentsBlock += tabulate([(lambda data: [
        data['flowId'][:12], # Flow
        str(data['total']), # Documents
        data['size(megabytes)'], # Size(MB)
        data['compressedSize(megabytes)'], # CompressedSize(MB)
        data['backupAt'].isoformat(timespec='milliseconds'), # BackupAt
        data['elapsed(s)'], # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'], # BackupSpeed(MB/s)
        data['total']/1440, # DocumentSpeed(doc/min)
])(_top10BiggestTotalFlows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

summary = '''```
Metathings Deviced Mongodb Backup Summary @{beginAt}
  Backup from: {beginAt} to: {endAt}
  Backup Job work from: {jobStartAt} to: {jobEndAt}, elapsed: {elapsedMinute:.2f} min
  Documents: {documents}, {documentsPerSecond:.2f} document/sec
  Size: {sizeMegabytes:.2f} MB, AverageDocumentSize: {avgDocumentSizeKilobytes:.2f} KB
  Compressed Size: {compressedSizeMegabytes:.2f} MB
{top10BiggestSizeDocumentsBlock}
{top10BiggestTotalDocumentsBlock}
```'''.format(**locals())

print(summary, flush=True)

requests.post(envSlackNotificationUrl, json={"text": summary})
