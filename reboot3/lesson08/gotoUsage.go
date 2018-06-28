package main

func main() {
	usage1()
	usage2()
	usage3()
}

func usage1() {
	// step 1 open file. handler
	if !ok {
		return
	}
	// step 2 allocate mem
	if !ok {
		return
		// handler 没释放，会泄露
	}
	// step 3 init
	if !ok {
		return
		// handler 没释放，会泄露
		// memory  没释放，会泄露
	}
	// step 4 Run work
}

func usage2() {
	// step 1 open file. handler of fd
	if !ok {
		return
	}
	// step 2 allocate mem fp
	if !ok {
		close(fd)
		return
	}
	// step 3 init
	if !ok {
		close(fd)
		free(fp)
		fp = NULL
		return
	}
	// step 4 Run work
}

func usage3() {
	// step 1 open file. handler of fd
	if !ok {
		return
	}
	// step 2 allocate mem fp
	if !ok {
		goto cleanFD
	}
	// step 3 init
	if !ok {
		goto cleanMEM
	}
	// step 4 Run work
	// no need to goto from here.

	// relax
cleanMEM:
	free(fp)
	fp = NULL
	// close(fd)
	// return
cleanFD:
	close(fd)
	return
}
