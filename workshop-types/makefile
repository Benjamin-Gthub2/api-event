documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd --instanceName workshop_types
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://127.0.0.1:8080/api/convert' -InFile 'docs\workshop_types_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_workshop_types_mocks:
	mockery --dir=domain --name=WorkshopTypeRepository --filename=workshop_types_repository_mock.go --output=domain/mocks --outpkg=workshop_types
	mockery --dir=domain --name=WorkshopTypeUseCase --filename=workshop_types_usecase_mock.go --output=domain/mocks --outpkg=workshop_types

PROJECT_PATH = ../../../

deploy-local-workshop_types:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/workshop_types.hr.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/workshop_types.hr.smartone:1.0.0
	rm -rf app
