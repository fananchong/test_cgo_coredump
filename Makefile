all:
	gcc -g3 -O0 -c test.c
	ar cr libtest.a test.o
	go build main.go

godump:
	./main > 1.log 2>&1

cdump:
	env GOTRACEBACK=crash ./main

clean:
	rm -f *.o main libtest.a core *.log
