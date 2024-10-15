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
