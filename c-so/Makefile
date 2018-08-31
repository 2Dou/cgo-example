c-so: main.go ./add/libadd.so
	go build && cp -f ./add/libadd.so ./

./add/libadd.so:
	cd ./add && make
	cd ..

clean:
	rm ./c-so ./*.so ./add/*.so
