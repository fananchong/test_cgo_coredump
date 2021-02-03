all:
	gcc -I. -g3 -O0 -c test.cpp catch_except.cpp
	ar cr libtest.a test.o catch_except.o
	go build -gcflags=all="-N -l" main.go
	go build -gcflags=all="-N -l" main2.go

godump:
	./main > 1.log 2>&1

cdump:
	env GOTRACEBACK=crash ./main

clean:
	rm -f *.o main libtest.a core *.log

cmain:
	gcc -I. -g3 -O0 test.cpp main.cpp -o main


