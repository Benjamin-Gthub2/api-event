documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd --instanceName sessions
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://127.0.0.1:8080/api/convert' -InFile 'docs\sessions_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_sessions_mocks:
	mockery --dir=domain --name=SessionRepository --filename=sessions_repository_mock.go --output=domain/mocks --outpkg=sessions
	mockery --dir=domain --name=SessionUseCase --filename=sessions_usecase_mock.go --output=domain/mocks --outpkg=sessions

PROJECT_PATH = ../../../

deploy-local-sessions:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/sessions.hr.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/sessions.hr.smartone:1.0.0
	rm -rf app
