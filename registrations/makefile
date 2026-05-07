documentation-win:
	swag init --pd --instanceName registrations
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://127.0.0.1:8080/api/convert' -InFile 'docs\registrations_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_registrations_mocks:
	mockery --dir=domain --name=RegistrationsRepository --filename=registrations_repository_mock.go --output=domain/mocks --outpkg=registrations
	mockery --dir=domain --name=RegistrationsUseCase --filename=registrations_usecase_mock.go --output=domain/mocks --outpkg=registrations

PROJECT_PATH = ../../../

deploy-local-warehouse-transfer-receipts:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/registrations.logistics.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/registrations.logistics.smartone:1.0.0
	rm -rf app
