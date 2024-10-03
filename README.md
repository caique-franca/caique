[transforms.redhat]
  type = "remap"
  inputs = ["filebeat"]
  source = '''
  .app = parse_regex(.log.file.path, r'^/home/Falcon/logs/falcon-(?P<app>[^/]+)/').app ??
         parse_regex(.log.file.path, r'^/data/dwc/app/.+-(?P<core>[^/]+)-(?P<app>[^/]+)/').app ??
         parse_regex(.log.file.path, r'^/data/dwyb/app/.+-(?P<core>[^/]+)-(?P<app>[^/]+)/').app ??
         parse_regex(.log.file.path, r'^/home/svc_dwclob/dwclob/(?P<app>[^/]+)/').app ??
         parse_regex(.log.file.path, r'^/data/dwc/(?P<app>[^/]+)/').app ??
         parse_regex(.log.file.path, r'^/data/dwyb/(?P<app>[^/]+)/').app
  '''
