# Tenta aplicar o regex e captura o erro
.parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}[/-]\d{2}[/-]\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')

# Se o primeiro regex falhar, tenta o próximo
if err != null {
  .parse, err = parse_regex(.message, r'(?P<event_timestamp>\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}):(?P<common_log_message>.*)')
}

if err != null {
  .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{2}/\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')
}

if err != null {
  .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')
}

if err != null {
  .parse, err = parse_regex(.message, r'(?P<event_timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z)\s*(?P<common_log_message>.*)')
}

if err != null {
  .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')
}

# Verifica se houve algum erro após todas as tentativas de regex
if err != null {
  .tags = ["regex_failure_beats_common_date"]
} else {
  # Define o timestamp do evento e deleta o campo temporário
  .event_timestamp = .parse.event_timestamp
  del(.parse.event_timestamp)
}

# Verifica se "env" e "region" estão faltando
if is_null(.fields.env) {
  .fields.env = "missing"
  .fields.error_info = "fields.env missing"
}

if is_null(.fields.region) {
  .fields.region = "missing"
  .fields.error_info = "fields.region missing"
}

# Atualiza o campo "error_info"
.fields.error_info = "Processed by " + .fields.region + "-" + .fields.env + " logstash"
