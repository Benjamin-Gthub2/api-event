documentation:
	swag init --pd --instanceName events
	curl -X POST http://192.168.71.200:8080/api/convert -d @docs/events_swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd --instanceName events
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://127.0.0.1:8080/api/convert' -InFile 'docs\events_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

events_mocks:
	mockery --dir=domain --name=EventRepository --filename=events_repository_mock.go --output=domain/mocks --outpkg=events
	mockery --dir=domain --name=EventUseCase --filename=events_usecase_mock.go --output=domain/mocks --outpkg=events


PROJECT_PATH = ../../../

deploy-local-events:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/events.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/events.core.smartone:1.0.0
	rm -rf app
