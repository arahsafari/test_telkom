soal1:
	@echo "Running soal 1..."
	go run ./no1/main.go

soal2:
	@echo "Running soal 2..."
	go run ./no2/main.go

soal5:
	@echo "Running soal 5..."
	docker-compose -f ./no5/docker-compose.yml up -d --build