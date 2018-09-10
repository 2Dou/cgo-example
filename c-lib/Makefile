c-lib: main.go ./add/libadd.a
	go build

./add/libadd.a:
	cd ./add && make
	cd ..

clean:
	rm ./c-lib ./add/*.a
