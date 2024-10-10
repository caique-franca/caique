type: "remap"
inputs:
  - _SYSTEM_filebeat
source: |
  .parse, err = parse_regex(.message, r'^(?P<event_timestamp>\d{4}[/-]\d{2}[/-]\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
                parse_regex(.message, r'(?P<event_timestamp>\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}):(?P<common_log_message>.*)') ??
                parse_regex(.message, r'^(?P<event_timestamp>\d{2}/\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
                parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)') ??
                parse_regex(.message, r'(?P<event_timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z)\s*(?P<common_log_message>.*)') ??
                parse_regex(.message, r'^(?P<event_timestamp>\d{4}\.\d{2}\.\d{2}\s*\d{2}:\d{2}:\d{2})\s*(?P<common_log_message>.*)')

  if err != null {
    .tags = "regex_failure_beats_common_date"
  }

  .event_timestamp = .parse.event_timestamp
  del(.parse)
  
  if is_null(.fields.env) {
    .fields.env = "missing"
    .fields.error_info = "fields.env missing"
  }

  if is_null(.fields.region) {
    .fields.region = "missing"
    .fields.error_info = "fields.region missing"
  }

  .fields.error_info, err = "Processed by " + .fields.region + "-" + .fields.env + " logstash"



message:[2024-10-10 10:11:46,297] [RTRESULT] [tomcat-http--26] [cwckywqmaerjrmi0xbbehqu3h7f dada 10.170.217.208] ServletUtil INFO getCookie(dada-si): sessionID=javax.servlet.http.Cookie@7ddf79aa
message:172.16.214.161 - - [10/Oct/2024:10:11:34 -0400] "POST /longpoll/listen.json HTTP/1.1" 200 439 "COOKIE[janney-si=9bee748dc015e7a8fd01ddb7a769d203a99a7968; tw-si-qa1=9bee748dc015e7a8fd01ddb7a769d203a99a7968; last-login-si=9bee748dc015e7a8fd01ddb7a769d203a99a7968; last-login-brand=janney; bd_auto_dest=0; TS0193c126=01faf79f112c4a7df27c8e10312b7cfc0b50f4969accb0b830aa5f196ff75f5a04c3d1a5af330f98c70abc4c7456462107eeb14fec39729b06044ed96bc5ab6c23301eea7780f94ce1473c7d00762a410803cf49bddffa2da4f8ecd3a0edbafffa7e48b5265850e7737beb99205fdbafe8e60eb074c27fa39c58eb66cf15cad665c90a0ccf; RT=1209472010.5163.0000]" [Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36] "REFERER[https://rt.qa.bonddesk.com/janney/servlet/longpoll/connect.html]" "BD_REQUEST_ID[d81d19733e1dcd173dd64f0fac865f6e]"
