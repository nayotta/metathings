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
df['compressionRatio'] = df['size(bytes)']/df['compressedSize(bytes)']
df['avgDocumentSize(KB/doc)'] = df['size(bytes)']/1024/df['total']

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

documentSize_Top10Desc_Block = '''
Document Size(top10, desc)
'''

_documentSize_Top10Desc_Flows = df.sort_values(by=['size(megabytes)'], ascending=False)
documentSize_Top10Desc_Block += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['compressionRatio'],  # CompressionRatio
        data['avgDocumentSize(KB/doc)'],  # AvgDocumentSize(KB/doc)
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_documentSize_Top10Desc_Flows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'CompressionRatio', 'AvgDocumentSize(KB/doc)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

documentTotal_Top10Desc_Block = '''
Document Total(top10, desc)
'''

_documentTotal_Top10Desc_Flows = df.sort_values(by=['total'], ascending=False)
documentTotal_Top10Desc_Block += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['compressionRatio'],  # CompressionRatio
        data['avgDocumentSize(KB/doc)'],  # AvgDocumentSize(KB/doc)    
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_documentTotal_Top10Desc_Flows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'CompressionRatio', 'AvgDocumentSize(KB/doc)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

documentAvgSize_Top10Desc_Block = '''
Document Avg Size(top10, desc)
'''

_documentAvgSize_Top10Desc_Flows = df.sort_values(by=['avgDocumentSize(KB/doc)'], ascending=False)
documentAvgSize_Top10Desc_Block += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['compressionRatio'],  # CompressionRatio
        data['avgDocumentSize(KB/doc)'],  # AvgDocumentSize(KB/doc)    
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_documentAvgSize_Top10Desc_Flows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'CompressionRatio', 'AvgDocumentSize(KB/doc)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

compressionRatio_Top10Desc_Block = '''
Compression Ratio(top10, desc)
'''

_compressionRatio_Top10Desc_Flows = df.sort_values(by=['compressionRatio'], ascending=False)
compressionRatio_Top10Desc_Block += tabulate([(lambda data: [
        data['flowId'][:12],  # Flow
        str(data['total']),  # Documents
        data['size(megabytes)'],  # Size(MB)
        data['compressedSize(megabytes)'],  # CompressedSize(MB)
        data['compressionRatio'],  # CompressionRatio
        data['avgDocumentSize(KB/doc)'],  # AvgDocumentSize(KB/doc)    
        data['backupAt'].isoformat(timespec='milliseconds'),  # BackupAt
        data['elapsed(s)'],  # Elapsed(s)
        data['size(megabytes)']/data['elapsed(s)'],  # BackupSpeed(MB/s)
        data['total']/1440,  # DocumentSpeed(doc/min)
])(_compressionRatio_Top10Desc_Flows.iloc[idx]) for idx in range(10)],
               headers=['Flow', 'Documents', 'Size(MB)', 'CompressedSize(MB)', 'CompressionRatio', 'AvgDocumentSize(KB/doc)', 'BackupAt', 'Elapsed(s)', 'BackupSpeed(MB/s)', 'DocumentSpeed(doc/min)'],
               numalign="right",
               floatfmt=".2f")

distributionsBlock = '''
Distributions
'''
quantileFactors = [1, 0.99, 0.95, 0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1, 0.05, 0.01, 0]
distributionsBlock += tabulate([(lambda name, total, mean, stddev, qs: [
    name,  # name
    total,  # total
    mean,  # mean
    stddev,  # stddev
    qs.iloc[0],  # Max
    qs.iloc[1],  # 99%
    qs.iloc[2],  # 95%
    qs.iloc[3],  # 90%
    qs.iloc[4],  # 80%
    qs.iloc[5],  # 70%
    qs.iloc[6],  # 60%
    qs.iloc[7],  # Mid
    qs.iloc[8],  # 40%
    qs.iloc[9],  # 30%
    qs.iloc[10],  # 20%
    qs.iloc[11],  # 10%
    qs.iloc[12],  # 5%
    qs.iloc[13],  # 1%
    qs.iloc[14],  # Min
])(*x) for x in [
    ['documents', len(df['total']), df['total'].mean(), df['total'].std(), df['total'].quantile(quantileFactors)],
    ['size(MB)', len(df['size(megabytes)']), df['size(megabytes)'].mean(), df['size(megabytes)'].std(), df['size(megabytes)'].quantile(quantileFactors)],
    ['avgSize(KB/doc)', len(df['avgDocumentSize(KB/doc)']), df['avgDocumentSize(KB/doc)'].mean(), df['avgDocumentSize(KB/doc)'].std(), df['avgDocumentSize(KB/doc)'].quantile(quantileFactors)],
    ['compressionRatio', len(df['compressionRatio']), df['compressionRatio'].mean(), df['compressionRatio'].std(), df['compressionRatio'].quantile(quantileFactors)],
]], headers=['', 'Total', 'Mean', 'Stddev', 'Max', '99%', '95%', '90%', '80%', '70%', '60%', 'Mid', '40%', '30%', '20%', '10%', '5%', '1%', 'Min'], numalign='right', floatfmt='.2f')

summaryBlock = '''Metathings Deviced Mongodb Backup Summary @{beginAt}
  Backup from: {beginAt} to: {endAt}
  Backup Job work from: {jobStartAt} to: {jobEndAt}, elapsed: {elapsedMinute:.2f} min
  Flows: {flows}, Documents: {documents}, {documentsPerSecond:.2f} document/sec, AverageDocuments: {avgDocuments:.2f}
  Size: {sizeMegabytes:.2f} MB, Compressed Size: {compressedSizeMegabytes:.2f} MB, AverageDocumentSize: {avgDocumentSizeKilobytes:.2f} KB
'''.format(**locals())

summary = '''```{summaryBlock}```'''.format(**locals())
documentSize_Top10Desc_Documents = '''```{documentSize_Top10Desc_Block}```'''.format(**locals())
documentTotal_Top10Desc_Documents = '''```{documentTotal_Top10Desc_Block}```'''.format(**locals())
documentAvgSize_Top10Desc_Documents = '''```{documentAvgSize_Top10Desc_Block}```'''.format(**locals())
compressionRatio_Top10Desc_Documents = '''```{compressionRatio_Top10Desc_Block}```'''.format(**locals())
distributions = '''```{distributionsBlock}```'''.format(**locals())

print(summary, flush=True)
print(documentSize_Top10Desc_Documents, flush=True)
print(documentTotal_Top10Desc_Documents, flush=True)
print(documentAvgSize_Top10Desc_Documents, flush=True)
print(compressionRatio_Top10Desc_Documents, flush=True)
print(distributions, flush=True)

requests.post(envSlackNotificationUrl, json={"text": summary})
requests.post(envSlackNotificationUrl, json={"text": documentSize_Top10Desc_Documents})
requests.post(envSlackNotificationUrl, json={"text": documentTotal_Top10Desc_Documents})
requests.post(envSlackNotificationUrl, json={"text": documentAvgSize_Top10Desc_Documents})
requests.post(envSlackNotificationUrl, json={"text": compressionRatio_Top10Desc_Documents})
requests.post(envSlackNotificationUrl, json={"text": distributions})
