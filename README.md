# lshw-xml-server

***⚠️ I developed this project to suit very specific needs. If you want to reuse it you'll probably need some coding to make it work.***

### Install

To install, you can either clone this project, get it using Go "get" command.

`git clone https://github.com/lucchmielowski/lshw-xml-saver.git`

or

`go get github.com/lucchmielowski/lshw-xml-saver`

### Usage

In the app directory:

`./lshw-xml-server -dir path_to_dir`

To work, the executable needs a folder with sub-folders containing xml results of the lshw command. **The sub-folder and the xml file it contains must always begin the same way**, and the xml file name should end with **"-ALL-XML"**

eg:
```
/my-folder
   /Server-1
      Serveur-1-ALL-XML.xml
   /Server-2
      Serveur-2-ALL-XML.xml
  ...
   /My-Server
      My-Server-ALL-XML.xml
```

### Command line parameters

```
  -db string
        Database name (default "lshw-xml")      
  -dir string
        Path to the directory containing all servers sub-directories (default "./files")
```



You can get the list of all parameters using `-h` tag
