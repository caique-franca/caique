ARG TW_BASE_IMAGE
FROM --platform=linux/arm64 ${TW_BASE_IMAGE}
LABEL tw.image=true

ARG GO_SEMREL_GITLAB_VERSION
ARG ARTIFACTORY_USER
ARG ARTIFACTORY_PASS

COPY cloud-ca.crt /etc/pki/ca-trust/source/anchors/
COPY web-corp-2.crt /etc/pki/ca-trust/source/anchors/
COPY web-corp-root-2.crt /etc/pki/ca-trust/source/anchors/
COPY pip.conf /etc/
COPY requirements.txt requirements.txt
RUN update-ca-trust

RUN curl -L -u $ARTIFACTORY_USER:$ARTIFACTORY_PASS -o /usr/local/bin/release https://artifactory.com/artifactory/ptt-docker-dev/juhani/go-semrel-gitlab--arm64/v${GO_SEMREL_GITLAB_VERSION}/artifacts/release \
  && chmod +x /usr/local/bin/release

RUN apt update -y \
  && apt install -y git-core unzip make diffutils jq curl gcc gcc-c++ patchelf ccache patch zlib-devel sqlite sqlite-devel tk-devel libffi-devel xz-devel ncurses ncurses-devel openldap openldap-devel wget sshpass \
  && apt clean all

# RUN pip3 --no-cache-dir install -r requirements.txt && rm -rf requirements.txt
