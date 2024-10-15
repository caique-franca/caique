# Parse "app" field from message or log.file.path
  if [fields][os_family] == "redhat" {
    if [process][name] {
      grok {
        match => {
          "[process][name]" => "-(?<app>[^-]+)$"
        }
      }
    }
    else {
      grok {
        match => {
          "[log][file][path]" => [
                                  "^/home/Falcon/logs/falcon-(?<app>[^/]+)/",
                                  "^/data/dwc/app/.+-(?<core>[^/]+)-(?<app>[^/]+)/",
                                  "^/home/svc_dwclob/dwclob/(?<app>[^/]+)/",
                                  "^/data/dwc/(?<app>[^/]+)/"
                                 ]
        }
      }
    }
  }
  else if [fields][os_family] == "windows" {
    grok {
      match => {
        "[log][file][path]" => "^d:\\localhost\\log\\(?<app>.+?)(\.\d{8})?\.err"
      }
    }
  }


# Check the OS family and proceed with parsing the "app" field accordingly
if .fields.os_family == "redhat" {
    # If "process.name" exists, extract the "app" from it using regex
    if exists(.process.name) {
        .app, err = parse_regex(.process.name, r'-(?P<app>[^-]+)$').app
    }
    # If "process.name" does not exist, try extracting "app" from "log.file.path"
    else {
        .app, err = parse_regex(.log.file.path, r'^/home/Falcon/logs/falcon-(?P<app>[^/]+)/').app ??
                    parse_regex(.log.file.path, r'^/data/dwc/app/.+-(?P<core>[^/]+)-(?P<app>[^/]+)/').app ??
                    parse_regex(.log.file.path, r'^/home/svc_dwclob/dwclob/(?P<app>[^/]+)/').app ??
                    parse_regex(.log.file.path, r'^/data/dwc/(?P<app>[^/]+)/').app
    }
}
# If the OS family is "windows", parse "app" from the Windows file path format
else if .fields.os_family == "windows" {
    .app, err = parse_regex(.log.file.path, r'^d:\\localhost\\log\\(?P<app>.+?)(\.\d{8})?\.err').app
}
