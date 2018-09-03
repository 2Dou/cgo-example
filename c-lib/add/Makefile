
libadd.a: add.o
	ar -rcs libadd.a add.o

add.o: src/add.c
	gcc -c -o add.o src/add.c

clean:
	rm *.o *.a
