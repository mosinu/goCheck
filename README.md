# goCheck

golang filesystem check to get a hash of files.


Tested on Linux and Windows 10

version 03:
	check md5sum/sha512 of all files recurisely in given path
	usage hashit <path>

		example usage: 
			*.nix: ./gocheck [sha512|md5] /var/www/
			Windows: gocheck.exe [sha512|md5] C:\Users\mosinu\Downloads
