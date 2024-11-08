# Go Load Balancer POC

HTTP 요청에 대한 로드 밸런싱을 처리하는 Go 기반의 프록시 서버입니다. 각 노드별로 RPM(Requests Per Minute)과 BPM(Bytes Per Minute) 속도 제한을 지원합니다.

## 주요 기능

- 다중 백엔드 노드 지원
- 노드별 독립적인 속도 제한
  - RPM (Requests Per Minute) 제한
  - BPM (Bytes Per Minute) 제한
- HTTP 프록시 기능


## Out of Scope
- 실시간 상태 모니터링
- 지수 백오프
- 노드 상태 체크 (Health Check)

## 프로젝트 구조

```
go-balancer/
├── cmd/                    # 실행 파일
│   └── main.go            # 메인 애플리케이션 진입점
├── config/                 # 설정 관련
│   ├── config.go          # 설정 처리 로직
│   └── config.yaml        # 설정 파일
├── deployments/           # 배포 관련
│   └── Dockerfile        # 도커 빌드 파일
├── internal/              # 내부 패키지
│   ├── proxy/            # 프록시 관련
│   │   ├── loadbalancer.go  # 로드 밸런서 구현
│   │   └── node.go          # 노드 구현
│   └── server/           # 서버 관련
│       └── handler.go    # HTTP 핸들러
├── pkg/                   # 공용 패키지
│   └── utils/            # 유틸리티 함수
│       └── utils.go
└── tests/                # 테스트 파일
    ├── proxy_test.go     # 프록시 테스트
    └── server_test.go    # 서버 테스트
```

## 설치 방법

```bash
# 저장소 클론
git clone https://github.com/violetaplum/go-balancer.git

# 의존성 설치
go mod download

# 빌드
go build -o proxy cmd/main.go
```

## 설정 방법

`config/config.yaml` 파일에서 다음 설정을 구성할 수 있습니다:

```yaml
nodes:
  - url: "https://api.openai.com/v1/models"
    max_bpm: 1000000  # 1MB/minute
    max_rpm: 100      # 100 requests/minute
  - url: "https://api.openai.com/v1/models"
    max_bpm: 1500000  # 1.5MB/minute
    max_rpm: 150      # 150 requests/minute
  - url: "https://api.openai.com/v1/models"
    max_bpm: 2000000  # 2MB/minute
    max_rpm: 200      # 200 requests/minute
port: "8080"         # 서버 포트
```

## 실행 방법

```bash
# 직접 실행
./proxy

# 도커로 실행
docker build -t go-balancer .
docker run -p 8080:8080 go-balancer
```

## API 엔드포인트

- `GET /status` - 현재 노드들의 상태 확인
- `ANY /*` - 모든 요청을 백엔드 서버로 프록시

## 테스트

```bash
# 전체 테스트 실행
go test ./...

# 특정 패키지 테스트
go test ./internal/proxy
go test ./internal/server
```

## Docker 지원

프로젝트는 Docker를 통한 배포를 지원합니다. Dockerfile이 제공되며 다음 명령으로 빌드하고 실행할 수 있습니다:

```bash
docker build -t go-balancer .
docker run -p 8080:8080 go-balancer
```
