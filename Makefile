TARGETS = \
	kafka-receiver \
	kafka-receiver-win64.exe \
	kafka-receiver-win32.exe \


all: ${TARGETS}

clean:
	touch ${TARGETS}
	rm    ${TARGETS}

mrproper: clean
	touch dummy~
	rm        *~

kafka-receiver: kafka-receiver.go
	go build kafka-receiver.go

kafka-receiver-win64.exe: kafka-receiver.go
	GOOS=windows GOARCH=amd64 go build -o kafka-receiver-win64.exe kafka-receiver.go

kafka-receiver-win32.exe: kafka-receiver.go
	GOOS=windows GOARCH=386 go build -o kafka-receiver-win32.exe kafka-receiver.go

