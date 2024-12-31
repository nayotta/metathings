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

flows = len(df)
beginAt = df.beginAt.iloc[0].strftime('%Y-%m-%d')
endAt = df.endAt.iloc[0].strftime('%Y-%m-%d')
jobStartAt = df.backupAt.min().isoformat(timespec='milliseconds')
jobEndAt = pd.to_datetime(df.backupAt.max().value+(df['elapsed(s)'].iloc[-1]*1000000000)+28800000000000).tz_localize('Asia/Shanghai').isoformat(timespec='milliseconds')
elapsedMinute = df['elapsed(s)'].sum()/60
documents = df['total'].sum()
avgDocuments = documents/flows
documentsPerSecond = df['total'].sum()/86400
sizeMegabytes = df['size(bytes)'].sum()/1024/1024
avgDocumentSizeKilobytes = df['size(bytes)'].sum()/1024/df['total'].sum()
compressedSizeMegabytes = df['compressedSize(bytes)'].sum()/1024/1024

top10BiggestSizeDocumentsBlock = '''
Top 10 Biggest Size Doucments
'''

_top10BiggestSizeFlows = df.sort_values(by=['size(megabytes)'], ascending=False)
top10BiggestSizeDocumentsBlock += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_top10BiggestSizeFlows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

top10BiggestTotalDocumentsBlock = '''
Top 10 Biggest Total Doucments
'''

_top10BiggestTotalFlows = df.sort_values(by=['total'], ascending=False)
top10BiggestTotalDocumentsBlock += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_top10BiggestTotalFlows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

distributionsBlock = '''
Distributions
'''

distributionsBlock += tabulate([(lambda name, data: [
    name,  # name
    data.iloc[0],  # Max
    data.iloc[1],  # 99%
    data.iloc[2],  # 95%
    data.iloc[3],  # 80%
    data.iloc[4],  # Mid
    data.iloc[4],  # 20%
    data.iloc[5],  # 5%
    data.iloc[6],  # 1%
    data.iloc[7],  # Min
])(x[0], x[1]) for x in [
    ['documents', df.total.quantile([1,0.99,0.95,0.8,0.5,0.2,0.05,0.01,0])],
    ['size(MB)', df['size(megabytes)'].quantile([1,0.99,0.95,0.8,0.5,0.2,0.05,0.01,0])],
]], headers=['', 'Max', '99%', '95%', '80%', 'Mid', '20%', '5%', '1%', 'Min'], numalign='right', floatfmt='.2f')

summaryBlock = '''Metathings Deviced Mongodb Backup Summary @{beginAt}
  Backup from: {beginAt} to: {endAt}
  Backup Job work from: {jobStartAt} to: {jobEndAt}, elapsed: {elapsedMinute:.2f} min
  Flows: {flows}, Documents: {documents}, {documentsPerSecond:.2f} document/sec, AverageDocuments: {avgDocuments}
  Size: {sizeMegabytes:.2f} MB, Compressed Size: {compressedSizeMegabytes:.2f} MB, AverageDocumentSize: {avgDocumentSizeKilobytes:.2f} KB
'''.format(**locals())

summary = '''```{summaryBlock}```'''.format(**locals())
top10BiggestSizeDocuments = '''```{top10BiggestSizeDocumentsBlock}```'''.format(**locals())
top10BiggestTotalDocuments = '''```{top10BiggestTotalDocumentsBlock}```'''.format(**locals())
distributions = '''```{distributionsBlock}```'''.format(**locals())

print(summary, flush=True)
print(top10BiggestSizeDocuments, flush=True)
print(top10BiggestTotalDocuments, flush=True)
print(distributions, flush=True)

requests.post(envSlackNotificationUrl, json={"text": summary})
requests.post(envSlackNotificationUrl, json={"text": top10BiggestSizeDocuments})
requests.post(envSlackNotificationUrl, json={"text": top10BiggestTotalDocuments})
requests.post(envSlackNotificationUrl, json={"text": distributions})
