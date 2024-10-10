# Tenta capturar o timestamp em diferentes formatos, incluindo colchetes e IP
.parse, err = parse_regex(.message, r'\[?(?P<event_timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}(?:[,\.]\d{3})?)\]?\s*(?P<common_log_message>.*)') ??
               parse_regex(.message, r'^(?P<ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s-\s\[(?P<event_timestamp>\d{2}/\w{3}/\d{4}:\d{2}:\d{2}:\d{2}\s-\d{4})\]\s*(?P<common_log_message>.*)') ??
               parse_regex(.message, r'^(?P<event_timestamp>\d{4}[/-]\d{2}[/-]\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
               parse_regex(.message, r'(?P<event_timestamp>\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}):(?P<common_log_message>.*)') ??
               parse_regex(.message, r'^(?P<event_timestamp>\d{2}/\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
               parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
               parse_regex(.message, r'(?P<event_timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z)\s*(?P<common_log_message>.*)')

# Se o regex falhar, adiciona a tag de falha
if err != null {
  .tags = ["regex_failure_beats_common_date"]
}

# Define o event_timestamp e remove o campo temporário
.event_timestamp = .parse.event_timestamp
del(.parse.event_timestamp)
