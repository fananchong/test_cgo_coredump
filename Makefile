all:
	gcc -I. -g3 -O0 -c test.cpp
	ar cr libtest.a test.o
	go build -gcflags="-l" main.go

godump:
	./main > 1.log 2>&1

cdump:
	env GOTRACEBACK=crash ./main

clean:
	rm -f *.o main libtest.a core *.log

cmain:
	gcc -I. -g3 -O0 test.cpp main.cpp -o main
