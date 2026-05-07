documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd --instanceName people
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://127.0.0.1:8080/api/convert' -InFile 'docs\people_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_people_mocks:
	mockery --dir=domain --name=PeopleRepository --filename=people_repository_mock.go --output=domain/mocks --outpkg=people
	mockery --dir=domain --name=PeopleUseCase --filename=people_usecase_mock.go --output=domain/mocks --outpkg=people

PROJECT_PATH = ../../../

deploy-local-people:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/people.hr.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/people.hr.smartone:1.0.0
	rm -rf app
