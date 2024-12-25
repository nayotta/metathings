#!/bin/bash

set -x

BACKUP_DIR="/mnt/backup"

mongosh $MONGODB_URI --eval 'db.getCollectionNames().forEach(x=>console.log(x))' | tee /tmp/colls.txt
BEGIN_AT=$(date -d '1 day ago' +%F)
END_AT=$(date -d '0 day ago' +%F)
BEGIN_AT_YEAR=$(echo $BEGIN_AT|cut -d '-' -f 1)
BEGIN_AT_MONTH=$(echo $BEGIN_AT|cut -d '-' -f 2)
BEGIN_AT_DAY=$(echo $BEGIN_AT|cut -d '-' -f 3)

BACKUP_DIR_WITH_DATE="$BACKUP_DIR/$BEGIN_AT_YEAR/$BEGIN_AT_MONTH/$BEGIN_AT_DAY"
mkdir -p ${BACKUP_DIR_WITH_DATE}

echo "beginAt,flowId,path,sha256sum,collection,total,size(byte),compressedSize(byte),backupAt,elapsed(s),endAt" | tee /tmp/metadata_v1.csv

for coll in $(cat /tmp/colls.txt); do
    flw=$(echo $coll|cut -d '.' -f2)
    full_path="$BACKUP_DIR_WITH_DATE/${flw:0:2}/${flw:2:2}/${flw:4:2}/${flw:6:99}"
    mkdir -p $full_path
    output_jsonl="output.jsonl"
    output_compressed_jsonl="${output_jsonl}.gz"
    tmp_jsonl="/tmp/${output_jsonl}"
    tmp_compressed_jsonl="/tmp/${output_compressed_jsonl}"
    backup_begin_at=$(date +%s%N)
    mongoexport $MONGODB_URI \
                --quiet \
                -d mt_devd_flw \
                -c $coll \
                --query '{"#dt": {"$gte": {"$date": "'$BEGIN_AT'T00:00:00+08:00"}, "$lt": {"$date": "'$END_AT'T00:00:00+08:00"}}}' \
                -o ${tmp_jsonl}
    backup_end_at=$(date +%s%N)
    backup_begin_at_rfc3339ns=$(date -d@$(printf "%d.%09d" $((backup_begin_at/1000000000)) $((backup_begin_at%1000000000))) --rfc-3339=ns)
    elapsed_ns=$(($backup_end_at-$backup_begin_at))
    elapsed_s=$(printf "%d.%03d" $((elapsed_ns/1000000000)) $((elapsed_ns%1000000000/1000000)))
    total=$(wc -l ${tmp_jsonl}|cut -d ' ' -f 1)
    size=$(stat -c%s ${tmp_jsonl})
    gzip $tmp_jsonl --stdout > ${tmp_compressed_jsonl}
    tmp_compressed_jsonl_sha256=$(sha256sum ${tmp_compressed_jsonl} | cut -d ' ' -f 1)
    compressed_size=$(stat -c%s ${tmp_compressed_jsonl})
    mv ${tmp_compressed_jsonl} ${full_path}
    echo "\"${BEGIN_AT}\",\"${flw}\",\"${full_path}/${output_compressed_jsonl}\",\"${tmp_compressed_jsonl_sha256}\",\"${coll}\",${total},${size},${compressed_size},\"${backup_begin_at_rfc3339ns}\",${elapsed_s},\"${END_AT}\"" | tee -a /tmp/metadata_v1.csv
done

mv /tmp/metadata_v1.csv ${BACKUP_DIR_WITH_DATE}
