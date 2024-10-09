  .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}[/-]\d{2}[/-]\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
             parse_regex(.message, r'(?P<event_timestamp>\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}):(?P<common_log_message>.*)') ??
             parse_regex(.message, r'^(?P<event_timestamp>\d{2}/\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
             parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
             parse_regex(.message, r'(?P<event_timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z)\s*(?P<common_log_message>.*)') ??
             parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')

  .event_timestamp = .parse.event_timestamp
  del(.parse.event_timestamp)
  if err != null {
    .tags = ["regex_failure_beats_common_date"]
  }
  if is_null(.fields.env) {
    .fields.env = "missing"
    .fields.error_info = "fields.env missing"
  }

  if is_null(.fields.region) {
    .fields.region = "missing"
    .fields.error_info = "fields.region missing"
  }

  .fields.error_info, err = "Processed by " + .fields.region + "-" + .fields.env + " logstash"


  2024-10-09T00:08:08.178478Z  WARN transform{component_kind="transform" component_id=twd_us-elk-dev_to_coralogix component_type=remap}: vector::transforms::remap: VRL compilation warning. warnings=
warning[E900]: unused variable `err`
  ┌─ :1:9
  │
1 │ .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}[/-]\d{2}[/-]\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
  │         ^^^ help: use the result of this expression or remove it
  │
  = this expression has no side-effects
  = see language documentation at https://vrl.dev
  = try your code in the VRL REPL, learn more at https://vrl.dev/examples