# txp 

Multi-platform golang Desktop GUI application for converting AEB 19.14 Direct Debit TXT file to SEPA XML (Core scheme)

This utility can parse a **AEB 19-14** direct debit text file, and generate a SEPA  ISO-20022 XML **PAIN.008.001.02** (Core scheme) direct debit file.

TXP uses [therecipe/qt](https://github.com/therecipe/qt) QT bindings for golang. Compiled releases available for Linux and Windows at [releases](releases)

![TXP APP](https://raw.githubusercontent.com/bercab/txp/master/screenshots/txp3.png)

*Motivation*: European banks no longer admit Direct Debit files in plain text AEB 19.14 format, and some legacy programs do not provide means to transform its output to the new European SEPA norm.

Interface in spanish language.

This tool uses sepakit golang module for parsing SEPA files. Please refer to [github.com/APSL/sepakit](https://github.com/APSL/sepakit) for more info.


## Build

This build process uses therecipe/qt builder with docker, so there is no need to install qt5 dev libraries.
More info at https://github.com/therecipe/qt/wiki/Deploying-Application

Install sepakit dependence on gopath:

```
go get -u github.com/apsl/sepakit
```

### Build for linux QT12 with docker:

```
docker pull therecipe/qt:linux
qtdeploy -docker build linux   
```
The executable is created under a deploy/ directory.

### Build for ubuntu 18.04 (QT 5.9.5)

```
sudo apt install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev
sudo apt install libqt*5-dev qt*5-dev qt*5-doc-html
export QT_PKG_CONFIG=true
go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...
qtsetup
qtdeploy build desktop
```

### Build on linux for Windows 64 static with docker:

```
docker pull therecipe/qt:windows_64_static
qtdeploy -docker build windows_64_static
qtdeploy -docker build linux   

```
The executable is created under a deploy/ directory.

## Author

Bernardo Cabezas - bcabezas at [apsl.net](https://www.apsl.net) - [bercab](https://github.com/bercab)

## License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details
