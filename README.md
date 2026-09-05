# Process Logs

Program for filtering JSONL logs by time range.

## Input file format

Each line is JSON:
```json
{"id": 12323, "timestamp": "2026/10/15-12:40:00+0300", "value": 235.4}
{"id": 123455, "timestamp": "2026/10/15-15:20:00+0300", "value": 453.3}
{"id": 12421312, "timestamp": "2026/10/15-16:10:00+0300", "value": 783.1}
```


## Output file
Each line is JSON, but the timestamp is in RFC3339 format and the entries are filtered by time from start to end:
```json
{"id":12323,"timestamp":"2026-10-15T09:40:00Z","value":235.4}
{"id":123455,"timestamp":"2026-10-15T12:20:00Z","value":453.3}
