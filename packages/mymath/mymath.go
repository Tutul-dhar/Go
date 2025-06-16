package mymath

func Average(xs []float64) float64 {
	sum := 0.0

	for _, v := range xs {
		sum += v 
	}

	return sum / float64(len(xs))
}

/*
	tutul@tutul-dhar-PC:~/Documents/Go$ mkdir packages
	tutul@tutul-dhar-PC:~/Documents/Go$ cd packages
	tutul@tutul-dhar-PC:~/Documents/Go/packages$ code .
	tutul@tutul-dhar-PC:~/Documents/Go/packages$ cd ..
	tutul@tutul-dhar-PC:~/Documents/Go$ code .
	tutul@tutul-dhar-PC:~/Documents/Go$ go mod init example.com/mypackage
	go: creating new go.mod: module example.com/mypackage
	go: to add module requirements and sums:
		go mod tidy
	tutul@tutul-dhar-PC:~/Documents/Go$ mkdir mymath
	tutul@tutul-dhar-PC:~/Documents/Go$ cd mymath
	tutul@tutul-dhar-PC:~/Documents/Go/mymath$ echo mymath.go
	mymath.go
	tutul@tutul-dhar-PC:~/Documents/Go/mymath$ cd ..
	tutul@tutul-dhar-PC:~/.local/share/Trash/files$ cd ..
	tutul@tutul-dhar-PC:~/.local/share/Trash$ cd ..
	tutul@tutul-dhar-PC:~/.local/share$ cd ..
	tutul@tutul-dhar-PC:~/.local$ cd ..
	tutul@tutul-dhar-PC:~$ cd Documents
	tutul@tutul-dhar-PC:~/Documents$ cd Go
	tutul@tutul-dhar-PC:~/Documents/Go$ mkdir packages
	tutul@tutul-dhar-PC:~/Documents/Go$ cd packages
	tutul@tutul-dhar-PC:~/Documents/Go/packages$ mkdir main
	tutul@tutul-dhar-PC:~/Documents/Go/packages$ mkdir mymath
*/