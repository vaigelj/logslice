# logslice

Fast log file slicer and filter utility with regex and time-range support.

---

## Installation

```bash
go install github.com/yourusername/logslice@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logslice.git && cd logslice && go build -o logslice .
```

---

## Usage

```
logslice [flags] <logfile>

Flags:
  -start string    Start time (e.g. "2024-01-15 08:00:00")
  -end string      End time   (e.g. "2024-01-15 09:00:00")
  -pattern string  Regex pattern to filter log lines
  -format string   Timestamp format (default: "2006-01-02 15:04:05")
  -o string        Output file (default: stdout)
```

### Examples

Filter logs by time range:
```bash
logslice -start "2024-01-15 08:00:00" -end "2024-01-15 09:00:00" app.log
```

Filter by regex pattern:
```bash
logslice -pattern "ERROR|WARN" app.log
```

Combine time range and pattern:
```bash
logslice -start "2024-01-15 08:00:00" -end "2024-01-15 09:00:00" -pattern "ERROR" -o errors.log app.log
```

---

## Features

- Fast line-by-line streaming — handles large log files with minimal memory usage
- Flexible time-range slicing with configurable timestamp formats
- Full regex support for pattern matching
- Output to file or stdout

---

## License

MIT © 2024 yourusername