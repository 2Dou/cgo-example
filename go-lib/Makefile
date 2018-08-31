go-lib: add.a
	gcc -o go-lib main.c add.a

add.a: ./add/add.go
	cd ./add && go build -buildmode=c-archive -o ../add.a
	cd ..

clean:
	rm go-lib *.a *.h
