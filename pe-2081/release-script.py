import os
import requests

pjid = os.environ.get('CI_PROJECT_ID')
jid = os.environ.get('CI_JOB_ID')
nv = os.environ.get('NEXT_VERSION')
gl = os.environ.get('CI_SERVER_HOST')
gt = os.environ.get('GITLAB_TOKEN')

assets = {
    "links": [
        {"name": "Linux amd64 binary", "url": f"https://{gl}/platform/forks/go-semrel-gitlab/-/jobs/{jid}/artifacts/file/build/linux_amd64/release"},
        {"name": "Darwin amd64 binary", "url": f"https://{gl}/platform/forks/go-semrel-gitlab/-/jobs/{jid}/artifacts/file/build/darwin_amd64/release"},
        {"name": "Windows amd64 binary", "url": f"https://{gl}/platform/forks/go-semrel-gitlab/-/jobs/{jid}/artifacts/file/build/windows_amd64/release"},
        {"name": "Linux arm64 binary", "url": f"https://{gl}/platform/forks/go-semrel-gitlab/-/jobs/{jid}/artifacts/file/build/linux_arm64/release"},
        {"name": "Darwin arm64 binary", "url": f"https://{gl}/platform/forks/go-semrel-gitlab/-/jobs/{jid}/artifacts/file/build/darwin_arm64/release"},
    ]
}

json = {
    "id": pjid,
    "name": nv,
    "tag_name": nv,
    "description": os.environ.get('CHANGELOG'),
    "ref": os.environ.get('HEAD_REF'),
    "assets": assets
}

print(f"POSTing release to GitLab: {json}")
resp = requests.post(f"https://{gl}/api/v4/projects/{pjid}/releases", json=json, headers={"PRIVATE-TOKEN": gt})
print(f"Received {resp.status_code} response: {resp.text}")
if resp.status_code < 200 or resp.status_code > 299:
  exit(1)
