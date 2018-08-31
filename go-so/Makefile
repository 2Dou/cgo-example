go-so: add.so
	gcc -o go-so main.c add.so

add.so: ./add/add.go
	cd ./add && go build -buildmode=c-shared -o ../add.so
	cd ..

clean:
	rm go-so *.so *.h
